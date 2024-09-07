package main

import (
	"fmt"

	"github.com/thebigyovadiaz/go_ether_blockchain/src/ether"
)

func main() {
	fmt.Println("Hello World - Ethereum Blockchain")

	balance := ether.GetBalanceAt()
	fmt.Printf("balance: %.2f ETH", balance)
}
