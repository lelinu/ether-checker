package ether_scan

import (
	"github.com/lelinu/api_utils/utils/error_utils"
	"github.com/lelinu/api_utils/utils/validator_utils"
)

type MultipleAddressBalancesRequest struct {
	Address string `json:"address"`
}

func (r *MultipleAddressBalancesRequest) Validate() *error_utils.ApiError {
	v := validator_utils.NewValidator()
	v.IsNotEmpty("address", r.Address)

	if !v.IsValid() {
		return error_utils.NewBadRequestError(v.Err.Error())
	}

	return nil
}

type MultipleAddressBalancesResponse struct {
	Result []Result `json:"result"`
}

