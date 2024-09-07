package ether

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

type credentials struct {
	name     string
	password string
}

func (c *credentials) setCredentials(name, password string) {
	if len(password) > 0 {
		c.password = password
	}

	if len(password) > 0 {
		c.password = password
	}
}

// generateEtherAccountAddr allow generated ethereum account address with phrase private
func generateEtherAccountAddr(cred credentials) string {
	key := keystore.NewKeyStore("../../wallet", keystore.StandardScryptN, keystore.StandardScryptP)

	credentials := credentials{}
	credentials.setCredentials(cred.name, cred.password)

	res, err := key.NewAccount(credentials.password)
	if err != nil {
		log.Fatal(err)
	}

	return res.Address.String()
}
