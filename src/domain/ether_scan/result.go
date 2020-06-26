package ether_scan

import (
	"math"
	"math/big"
)

type Result struct {
	Account string `json:"account"`
	Balance string `json:"balance"`
}

func (r *Result) GetEtherValue() *big.Float {
	balance := new(big.Float)
	balance.SetString(r.Balance)
	return new(big.Float).Quo(balance, big.NewFloat(math.Pow10(18)))
}



