package ether

import (
	"context"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// GetBalanceAt allow to get balance in an account
func GetBalanceAt() *big.Float {
	address := common.HexToAddress(ganacheAddrAccount)
	//blockNumber := getBlockByNumber()

	balanceWei, err := client().BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Error to get balance: %w", err)
	}

	balanceFloat := new(big.Float)
	balanceFloat.SetString(balanceWei.String())

	etherBalance := new(big.Float).Quo(balanceFloat, big.NewFloat(math.Pow10(18)))
	return etherBalance
}
