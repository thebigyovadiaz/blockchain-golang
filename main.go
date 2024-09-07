package main

import (
	"fmt"

	"github.com/thebigyovadiaz/go_ether_blockchain/src/ether"
)

func main() {
	fmt.Printf("Hello World - Ethereum Blockchain\n\n")

	// Last block number
	block := ether.GetBlockByNumber()
	fmt.Printf("block: %v\n\n", block)

	// Get balance account
	balance := ether.GetBalanceAt()
	fmt.Printf("balance: %.2f ETH\n\n", balance)

	// Get public address
	pubKeyAddr := ether.GeneratePublicAddress()
	fmt.Printf("pubKeyAddr: %s\n\n", pubKeyAddr)

	// Get ethereum account address
	etherAddr := ether.GenerateEtherAddress()
	fmt.Println("etherAddr: ", etherAddr)
}
