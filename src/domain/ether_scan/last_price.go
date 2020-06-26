package ether_scan

import (
	"math/big"
	"strconv"
)

type LastPriceRequest struct {

}

type LastPriceResponse struct {
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	LastPrice LastPrice `json:"result"`
}

type LastPrice struct {
	Usd string `json:"ethusd"`
}

func (l *LastPriceResponse) GetPrice() *big.Float{
	lPrice, _ := strconv.ParseFloat(l.LastPrice.Usd, 64)
	return big.NewFloat(lPrice)
}
