package ether_scan

import (
	"github.com/lelinu/api_utils/utils/error_utils"
	"github.com/lelinu/api_utils/utils/validator_utils"
)

type SingleAddressBalanceRequest struct {
	Address string `json:"address"`
}

func (r *SingleAddressBalanceRequest) Validate() *error_utils.ApiError {
	v := validator_utils.NewValidator()
	v.IsNotEmpty("address", r.Address)

	if !v.IsValid() {
		return error_utils.NewBadRequestError(v.Err.Error())
	}

	return nil
}

type SingleAddressBalance struct {
	Result string `json:"result"`
}

type SingleAddressBalanceResponse struct {
	Result *Result `json:"result"`
}

func NewSingleAddressBalanceResponse(address string, result string) *SingleAddressBalanceResponse{
	return &SingleAddressBalanceResponse{
		Result: &Result{
			Account: address,
			Balance: result,
		},
	}
}
