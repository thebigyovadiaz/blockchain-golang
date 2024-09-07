package ether

import (
	"context"
	"log"
	"math/big"
)

func getBlockByNumber() *big.Int {
	block, err := client().BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get a block: %w", err)
	}

	return block.Number()
}
