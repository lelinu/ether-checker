package ether_scan_provider

import (
	"encoding/json"
	"fmt"
	"github.com/lelinu/api_utils/utils/error_utils"
	"github.com/lelinu/ether-checker/src/domain/ether_scan"
	"github.com/lelinu/golang-restclient/rest"
	"strings"
	"time"
)

type provider struct {
	client  rest.RequestBuilder
	baseUrl string
	apiKey  string
}

type IProvider interface {
	GetSingleAddressBalance(req *ether_scan.SingleAddressBalanceRequest) (*ether_scan.SingleAddressBalanceResponse, *error_utils.ApiError)
	GetMultipleAddressBalances(req *ether_scan.MultipleAddressBalancesRequest) (*ether_scan.MultipleAddressBalancesResponse, *error_utils.ApiError)
	GetLastPrice(req *ether_scan.LastPriceRequest) (*ether_scan.LastPriceResponse, *error_utils.ApiError)
}

var (
	GetSingleAccountBalanceUrl    = "?module=account&action=balance&address=#address#&tag=latest"
	GetMultipleAccountBalancesUrl = "?module=account&action=balancemulti&address=#address#&tag=latest"
	GetLastPriceUrl = "?module=stats&action=ethprice"
)

func NewProvider(baseUrl string, apiKey string) *provider {

	return &provider{
		client: rest.RequestBuilder{
			BaseURL: baseUrl,
			Timeout: 1 * time.Second,
		},
		baseUrl: baseUrl,
		apiKey: apiKey}
}

func (es *provider) getUrl(pointUrl string, key string, value string) string {

	pointUrl = strings.Replace(pointUrl, key, strings.TrimSpace(value), -1)
	return fmt.Sprintf("%v&apikey=%v", pointUrl, es.apiKey)
}

func (es *provider) GetSingleAddressBalance(req *ether_scan.SingleAddressBalanceRequest) (*ether_scan.SingleAddressBalanceResponse, *error_utils.ApiError) {

	if apiErr := req.Validate(); apiErr != nil{
		return nil, apiErr
	}

	httpResponse := es.client.Get(es.getUrl(GetSingleAccountBalanceUrl, "#address#", req.Address))

	// timeout
	if httpResponse == nil || httpResponse.Response == nil {
		return nil, error_utils.NewInternalServerError("Invalid Rest Client while trying to get single account balance")
	}

	// invalid response code
	if httpResponse.StatusCode > 299 {
		var errResp ether_scan.Response
		if err := json.Unmarshal(httpResponse.Bytes(), &errResp); err != nil {
			return nil, error_utils.NewInternalServerError("Invalid Error Json Body while trying to get single account balance")
		}

		return nil, error_utils.NewInternalServerError(errResp.Message)
	}

	// get base response
	var baseResponse ether_scan.Response
	if err := json.Unmarshal(httpResponse.Bytes(), &baseResponse); err != nil {
		return nil, error_utils.NewInternalServerError("Invalid Error Json Body while trying to get single account balance")
	}

	// check if it's an error
	if baseResponse.IsError(){
		return nil, error_utils.NewInternalServerError(baseResponse.Message)
	}

	// read response again
	var response ether_scan.SingleAddressBalance
	if err := json.Unmarshal(httpResponse.Bytes(), &response); err != nil {
		fmt.Printf("error is %v", err)
		return nil, error_utils.NewInternalServerError("Invalid Json Body while trying to get single account balance")
	}

	return ether_scan.NewSingleAddressBalanceResponse(req.Address, response.Result), nil
}

func (es *provider) GetMultipleAddressBalances(req *ether_scan.MultipleAddressBalancesRequest) (*ether_scan.MultipleAddressBalancesResponse, *error_utils.ApiError) {

	if apiErr := req.Validate(); apiErr != nil{
		return nil, apiErr
	}

	httpResponse := es.client.Get(es.getUrl(GetMultipleAccountBalancesUrl, "#address#", req.Address))

	// timeout
	if httpResponse == nil || httpResponse.Response == nil {
		return nil, error_utils.NewInternalServerError("Invalid Rest Client while trying to get multiple account balances")
	}

	// invalid response code
	if httpResponse.StatusCode > 299 {
		var errResp ether_scan.Response
		if err := json.Unmarshal(httpResponse.Bytes(), &errResp); err != nil {
			return nil, error_utils.NewInternalServerError("Invalid Error Json Body while trying to get multiple account balances")
		}

		return nil, error_utils.NewInternalServerError(errResp.Message)
	}

	// get base response
	var baseResponse ether_scan.Response
	if err := json.Unmarshal(httpResponse.Bytes(), &baseResponse); err != nil {
		return nil, error_utils.NewInternalServerError("Invalid Error Json Body while trying to get  multiple account balances")
	}

	// check if it's an error
	if baseResponse.IsError(){
		return nil, error_utils.NewInternalServerError(baseResponse.Message)
	}

	// read response again
	var response ether_scan.MultipleAddressBalancesResponse
	if err := json.Unmarshal(httpResponse.Bytes(), &response); err != nil {
		return nil, error_utils.NewInternalServerError("Invalid Json Body while trying to get multiple account balances")
	}

	return &response, nil
}

func (es *provider) GetLastPrice(req *ether_scan.LastPriceRequest) (*ether_scan.LastPriceResponse, *error_utils.ApiError) {

	httpResponse := es.client.Get(es.getUrl(GetLastPriceUrl, "", ""))

	// timeout
	if httpResponse == nil || httpResponse.Response == nil {
		return nil, error_utils.NewInternalServerError("Invalid Rest Client while trying to get last price")
	}

	// invalid response code
	if httpResponse.StatusCode > 299 {
		var errResp ether_scan.Response
		if err := json.Unmarshal(httpResponse.Bytes(), &errResp); err != nil {
			return nil, error_utils.NewInternalServerError("Invalid Error Json Body while trying to get last price")
		}

		return nil, error_utils.NewInternalServerError(errResp.Message)
	}

	// get base response
	var baseResponse ether_scan.Response
	if err := json.Unmarshal(httpResponse.Bytes(), &baseResponse); err != nil {
		return nil, error_utils.NewInternalServerError("Invalid Error Json Body while trying to get last price")
	}

	// check if it's an error
	if baseResponse.IsError(){
		return nil, error_utils.NewInternalServerError(baseResponse.Message)
	}

	// read response again
	var response ether_scan.LastPriceResponse
	if err := json.Unmarshal(httpResponse.Bytes(), &response); err != nil {
		return nil, error_utils.NewInternalServerError("Invalid Json Body while trying to get last price")
	}

	return &response, nil
}











