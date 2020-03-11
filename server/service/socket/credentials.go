package socket

import "os"

type Credentials struct {
	Address string
	Realm   string
	User    string
	Pword   string
}

func GetCredentials() Credentials {
	return Credentials{
		Address: os.Getenv("NEXUS_ADDRESS"),
		Realm:   os.Getenv("NEXUS_REALM"),
		User:    os.Getenv("NEXUS_USER"),
		Pword:   os.Getenv("NEXUS_PWORD"),
	}
}
