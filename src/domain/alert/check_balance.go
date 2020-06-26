package alert

import (
	"github.com/lelinu/api_utils/utils/error_utils"
	"github.com/lelinu/api_utils/utils/validator_utils"
	"math/big"
)

type CheckBalanceRequest struct {
	Address     string    `json:"address"`
	MinUsdValue *big.Float `json:"min_usd_value"`
}

func (r *CheckBalanceRequest) Validate() *error_utils.ApiError {

	v := validator_utils.NewValidator()
	v.IsNotEmpty("address", r.Address)

	if !v.IsValid() {
		return error_utils.NewBadRequestError(v.Err.Error())
	}

	return nil
}
