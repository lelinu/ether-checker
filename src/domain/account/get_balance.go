package account

import (
	"github.com/lelinu/api_utils/utils/error_utils"
	"github.com/lelinu/api_utils/utils/validator_utils"
	"math/big"
)

type GetBalanceRequest struct {
	Address string `json:"address"`
}

func (r *GetBalanceRequest) Validate() *error_utils.ApiError {

	v := validator_utils.NewValidator()
	v.IsNotEmpty("address", r.Address)

	if !v.IsValid() {
		return error_utils.NewBadRequestError(v.Err.Error())
	}

	return nil
}

type GetBalanceResponse struct {
	Accounts []*Account `json:"account"`
}

func NewGetBalanceResponse(accounts []*Account) *GetBalanceResponse {
	return &GetBalanceResponse{Accounts: accounts}
}

type Account struct {
	Address    string     `json:"address"`
	EthBalance *big.Float `json:"eth_balance"`
	UsdValue   *big.Float `json:"usd_value"`
}

func NewAccount(address string, ethBalance *big.Float, lastPrice *big.Float) *Account {

	// calculate usd value
	usdValue := new(big.Float).Mul(ethBalance, lastPrice)

	return &Account{
		Address:    address,
		EthBalance: ethBalance,
		UsdValue:   usdValue,
	}
}
