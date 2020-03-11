package db

import "os"

type Credentials struct {
	Address    string
	Name       string
	User       string
	Pword      string
	Connection string
}

func GetCredentials() Credentials {
	address := os.Getenv("DB_ADDRESS")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	pword := os.Getenv("DB_PWORD")
	return Credentials{
		Address:    address,
		Name:       name,
		User:       user,
		Pword:      pword,
		Connection: conn(address, name, user, pword),
	}
}

func conn(address, name, user, pword string) string {
	return user + ":" + pword + "@tcp(" + address + ")/" + name + "?charset=utf8mb4,utf8&parseTime=True"
}
