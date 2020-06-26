package services

import (
	"github.com/lelinu/api_utils/log/lzap"
	"github.com/lelinu/api_utils/utils/error_utils"
	"github.com/lelinu/ether-checker/src/domain/account"
	"github.com/lelinu/ether-checker/src/domain/ether_scan"
	"github.com/lelinu/ether-checker/src/providers/ether_scan_provider"
)

type IAccountService interface {
	GetBalance(req *account.GetBalanceRequest) (*account.GetBalanceResponse, *error_utils.ApiError)
}

type accountService struct {
	logger      lzap.IService
	ethProvider ether_scan_provider.IProvider
}

func NewAccountService(logger lzap.IService, ethProvider ether_scan_provider.IProvider) IAccountService {
	return &accountService{logger: logger, ethProvider: ethProvider}
}

func (s *accountService) GetBalance(req *account.GetBalanceRequest) (*account.GetBalanceResponse, *error_utils.ApiError) {

	if err := req.Validate(); err != nil{
		return nil, err
	}

	// get balances
	resp, err := s.ethProvider.GetMultipleAddressBalances(&ether_scan.MultipleAddressBalancesRequest{Address: req.Address})
	if err != nil{
		return nil, err
	}

	// get last price
	lastPrice, err := s.ethProvider.GetLastPrice(nil)
	if err != nil{
		return nil, err
	}

	// build response
	accounts := make([]*account.Account, 0)
	for _, v := range resp.Result {

		account := account.NewAccount(v.Account, v.GetEtherValue(), lastPrice.GetPrice())
		accounts = append(accounts, account)
	}

	// return response
	return account.NewGetBalanceResponse(accounts), nil
}
