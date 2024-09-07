package ether

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	infuraURL          = "https://mainnet.infura.io/v3/846d117036d548b1b92e6a21a950cc0b"
	ganacheURL         = "http://127.0.0.1:7545"
	ganacheAddrAccount = "0x4D453c1cb6Cf42485b86afF0324a1D6f947316B9"
)

func client() *ethclient.Client {
	client, err := ethclient.DialContext(context.Background(), ganacheURL)
	if err != nil {
		log.Fatalf("Error to create a ether client: %w", err)
	}
	defer client.Close()

	return client
}
