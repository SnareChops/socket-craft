package db

type Migrations struct {
	Table    string
	Items    []Migration
	Existing []string
}

func InitMigrations(table string, migrations []MigrationFunc) *Migrations {
	ms := []Migration{}
	for _, fn := range migrations {
		ms = append(ms, InitMigration(fn))
	}
	return &Migrations{
		Table:    table,
		Items:    ms,
		Existing: []string{},
	}
}

func (m *Migrations) Run(client Client) {
	println("Running migrations...")
	client.MustExec("CREATE TABLE IF NOT EXISTS `" + m.Table + "` (`name` varchar(255), `timestamp` datetime, PRIMARY KEY(`name`))")
	m.Fetch(client)
	for _, migration := range m.Items {
		if !m.Exists(migration.Name()) {
			migration.Run(m.Table, client)
		}
	}
}

func (m *Migrations) Fetch(client Client) {
	if err := client.Select(&m.Existing, "SELECT `name` FROM `"+m.Table+"`"); err != nil {
		panic(err)
	}
}

func (m *Migrations) Exists(name string) bool {
	for _, item := range m.Existing {
		if item == name {
			return true
		}
	}
	return false
}
