package socket

type Model struct {
	Credentials   Credentials
	Actions       Actions
	Subscriptions Subscriptions
}

func InitModel(actions Actions, subscriptions Subscriptions) *Model {
	return &Model{
		Credentials:   GetCredentials(),
		Actions:       actions,
		Subscriptions: subscriptions,
	}
}
