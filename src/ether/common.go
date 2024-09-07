package ether

import "github.com/ethereum/go-ethereum/common/hexutil"

func encodeHex(key []byte) string {
	return hexutil.Encode(key)
}
