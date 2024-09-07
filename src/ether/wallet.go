package ether

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

// generatePublicKey allow generate a new public key
func generatePublicKey() ecdsa.PublicKey {
	pvKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	pvKeyData := crypto.FromECDSA(pvKey)
	fmt.Println("pvKeyData: ", encodeHex(pvKeyData))

	pubData := crypto.FromECDSAPub(&pvKey.PublicKey)
	fmt.Println("pubData: ", encodeHex(pubData))

	return pvKey.PublicKey
}

// GeneratePublicAddress allow generate a new public address
func GeneratePublicAddress() (pubKeyAddress string) {
	pubKey := generatePublicKey()
	pubKeyAddress = crypto.PubkeyToAddress(pubKey).Hex()
	return
}

// GenerateEtherAddress allow generate a new ethereum account address
func GenerateEtherAddress() string {
	cred := credentials{
		name: "Pepe Pepeto",
		password: "PhraseUpperSecureEther2024",
	}

	etherAddr := generateEtherAccountAddr(cred)
	return etherAddr
}
