package ether_scan_provider

import (
	"github.com/lelinu/ether-checker/src/domain/ether_scan"
	"github.com/lelinu/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

var (
	baseUrl = "https://test.com"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestGetUrlSuccessful(t *testing.T){
	provider := NewProvider(baseUrl, "")
	value := provider.getUrl("?module=account&action=balance&address=#address#&tag=latest", "#address#", "0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a")
	assert.EqualValues(t,"?module=account&action=balance&address=0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a&tag=latest&apikey=", value)
}

func TestGetSingleAddressBalanceErrorInvalidJsonBodyStatusNotOk(t *testing.T){
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:         baseUrl + "?module=account&action=balance&address=12345&tag=latest&apikey=",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusInternalServerError,
		RespBody:     `{"message": false }`,
	})

	provider := NewProvider(baseUrl, "")
	resp, err := provider.GetSingleAddressBalance(&ether_scan.SingleAddressBalanceRequest{Address: "12345"})
	assert.NotNil(t, err)
	assert.Nil(t, resp)
	assert.EqualValues(t, http.StatusInternalServerError, err.HttpStatusCode)
	assert.EqualValues(t, "Invalid Error Json Body while trying to get single account balance", err.ErrorMessage)
}

func TestGetSingleAddressBalanceErrorStatusNotOk(t *testing.T){
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:         baseUrl + "?module=account&action=balance&address=12345&tag=latest&apikey=",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusInternalServerError,
		RespBody:     `{
			"status": "0",
			"message": "NOTOK",
			"result": "0"
		}`,
	})

	provider := NewProvider(baseUrl, "")
	resp, err := provider.GetSingleAddressBalance(&ether_scan.SingleAddressBalanceRequest{ Address: "12345"})
	assert.NotNil(t, err)
	assert.Nil(t, resp)
	assert.EqualValues(t, http.StatusInternalServerError, err.HttpStatusCode)
	assert.EqualValues(t, "NOTOK", err.ErrorMessage)
}

func TestGetSingleAddressBalanceErrorInvalidJsonBodyStatusOk(t *testing.T){
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:         baseUrl + "?module=account&action=balance&address=12345&tag=latest&apikey=",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"message": false }`,
	})

	provider := NewProvider(baseUrl, "")
	resp, err := provider.GetSingleAddressBalance(&ether_scan.SingleAddressBalanceRequest{ Address: "12345" })
	assert.NotNil(t, err)
	assert.Nil(t, resp)
	assert.EqualValues(t, http.StatusInternalServerError, err.HttpStatusCode)
	assert.EqualValues(t, "Invalid Error Json Body while trying to get single account balance", err.ErrorMessage)
}

func TestGetSingleAddressBalanceErrorStatusOk(t *testing.T){
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:         baseUrl + "?module=account&action=balance&address=12345&tag=latest&apikey=",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{
			"status": "0",
			"message": "NOTOK",
			"result": "0"
		}`,
	})

	provider := NewProvider(baseUrl, "")
	resp, err := provider.GetSingleAddressBalance(&ether_scan.SingleAddressBalanceRequest{ Address: "12345" })
	assert.NotNil(t, err)
	assert.Nil(t, resp)
	assert.EqualValues(t, http.StatusInternalServerError, err.HttpStatusCode)
	assert.EqualValues(t, "NOTOK", err.ErrorMessage)
}

func TestGetSingleAddressBalanceSuccessful(t *testing.T){
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:         baseUrl + "?module=account&action=balance&address=0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a&tag=latest&apikey=",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{
			"status": "1",
			"message": "OK",
			"result": "40891631566070000000000"
		}`,
	})

	provider := NewProvider(baseUrl, "")
	resp, err := provider.GetSingleAddressBalance(&ether_scan.SingleAddressBalanceRequest{ Address: "0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a" })
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Result)
	assert.EqualValues(t, "40891631566070000000000", resp.Result.Balance)
	assert.EqualValues(t, "0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a", resp.Result.Account)
}

func TestGetMultipleAddressBalancesErrorInvalidJsonBodyStatusNotOk(t *testing.T){
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:         baseUrl + "?module=account&action=balancemulti&address=12345&tag=latest&apikey=",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusInternalServerError,
		RespBody:     `{"message": false }`,
	})

	provider := NewProvider(baseUrl, "")
	resp, err := provider.GetMultipleAddressBalances(&ether_scan.MultipleAddressBalancesRequest{ Address:"12345" })
	assert.NotNil(t, err)
	assert.Nil(t, resp)
	assert.EqualValues(t, http.StatusInternalServerError, err.HttpStatusCode)
	assert.EqualValues(t, "Invalid Error Json Body while trying to get multiple account balances", err.ErrorMessage)
}

func TestGetMultipleAddressBalancesErrorStatusNotOk(t *testing.T){
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:         baseUrl + "?module=account&action=balancemulti&address=12345&tag=latest&apikey=",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusInternalServerError,
		RespBody:     `{
			"status": "0",
			"message": "NOTOK",
			"result": "0"
		}`,
	})

	provider := NewProvider(baseUrl, "")
	resp, err := provider.GetMultipleAddressBalances(&ether_scan.MultipleAddressBalancesRequest{Address:"12345"})
	assert.NotNil(t, err)
	assert.Nil(t, resp)
	assert.EqualValues(t, http.StatusInternalServerError, err.HttpStatusCode)
	assert.EqualValues(t, "NOTOK", err.ErrorMessage)
}

func TestGetMultipleAddressBalanceErrorInvalidJsonBodyStatusOk(t *testing.T){
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:         baseUrl + "?module=account&action=balancemulti&address=12345&tag=latest&apikey=",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"message": false }`,
	})

	provider := NewProvider(baseUrl, "")
	resp, err := provider.GetMultipleAddressBalances(&ether_scan.MultipleAddressBalancesRequest{Address: "12345" })
	assert.NotNil(t, err)
	assert.Nil(t, resp)
	assert.EqualValues(t, http.StatusInternalServerError, err.HttpStatusCode)
	assert.EqualValues(t, "Invalid Error Json Body while trying to get  multiple account balances", err.ErrorMessage)
}

func TestGetMultipleAddressBalancesErrorStatusOk(t *testing.T){
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:         baseUrl + "?module=account&action=balancemulti&address=12345&tag=latest&apikey=",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{
			"status": "0",
			"message": "NOTOK",
			"result": "0"
		}`,
	})

	provider := NewProvider(baseUrl, "")
	resp, err := provider.GetMultipleAddressBalances(&ether_scan.MultipleAddressBalancesRequest{Address: "12345"})
	assert.NotNil(t, err)
	assert.Nil(t, resp)
	assert.EqualValues(t, http.StatusInternalServerError, err.HttpStatusCode)
	assert.EqualValues(t, "NOTOK", err.ErrorMessage)
}

func TestGetMultipleAddressBalancesSuccessful(t *testing.T){
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:         baseUrl + "?module=account&action=balancemulti&address=0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a,0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a&tag=latest&apikey=",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{
			"status": "1",
			"message": "OK",
			"result": [
				{
					"account": "0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a",
					"balance": "40891631566070000000000"
				},
				{
					"account": "0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a",
					"balance": "40891631566070000000000"
				}
			]
		}`,
	})

	provider := NewProvider(baseUrl, "")
	resp, err := provider.GetMultipleAddressBalances(&ether_scan.MultipleAddressBalancesRequest{ Address: "0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a,0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a" })
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Result)
	assert.EqualValues(t, 2, len(resp.Result))
	assert.EqualValues(t, "0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a", resp.Result[0].Account)
	assert.EqualValues(t, "40891631566070000000000", resp.Result[0].Balance)
	assert.EqualValues(t, "0xddbd2b932c763ba5b1b7ae3b362eac3e8d40121a", resp.Result[1].Account)
	assert.EqualValues(t, "40891631566070000000000", resp.Result[1].Balance)
}

func TestGetLastPriceErrorInvalidJsonBodyStatusNotOk(t *testing.T){
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:         baseUrl + "?module=stats&action=ethprice&apikey=",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusInternalServerError,
		RespBody:     `{"message": false }`,
	})

	provider := NewProvider(baseUrl, "")
	resp, err := provider.GetLastPrice(nil)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
	assert.EqualValues(t, http.StatusInternalServerError, err.HttpStatusCode)
	assert.EqualValues(t, "Invalid Error Json Body while trying to get last price", err.ErrorMessage)
}

func TestGetLastPriceErrorStatusNotOk(t *testing.T){
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:         baseUrl + "?module=stats&action=ethprice&apikey=",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusInternalServerError,
		RespBody:     `{
			"status": "0",
			"message": "NOTOK",
			"result": "0"
		}`,
	})

	provider := NewProvider(baseUrl, "")
	resp, err := provider.GetLastPrice(nil)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
	assert.EqualValues(t, http.StatusInternalServerError, err.HttpStatusCode)
	assert.EqualValues(t, "NOTOK", err.ErrorMessage)
}

func TestGetLastPriceErrorInvalidJsonBodyStatusOk(t *testing.T){
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:         baseUrl + "?module=stats&action=ethprice&apikey=",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"message": false }`,
	})

	provider := NewProvider(baseUrl, "")
	resp, err := provider.GetLastPrice(nil)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
	assert.EqualValues(t, http.StatusInternalServerError, err.HttpStatusCode)
	assert.EqualValues(t, "Invalid Error Json Body while trying to get last price", err.ErrorMessage)
}

func TestGetLastPriceErrorStatusOk(t *testing.T){
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:         baseUrl + "?module=stats&action=ethprice&apikey=",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{
			"status": "0",
			"message": "NOTOK",
			"result": "0"
		}`,
	})

	provider := NewProvider(baseUrl, "")
	resp, err := provider.GetLastPrice(nil)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
	assert.EqualValues(t, http.StatusInternalServerError, err.HttpStatusCode)
	assert.EqualValues(t, "NOTOK", err.ErrorMessage)
}

func TestGetLastPriceSuccessful(t *testing.T){
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          baseUrl + "?module=stats&action=ethprice&apikey=",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"status":"1","message":"OK","result":{"ethbtc":"0.02523","ethbtc_timestamp":"1593103495","ethusd":"234.22","ethusd_timestamp":"1593103494"}}`,
	})

	provider := NewProvider(baseUrl, "")
	resp, err := provider.GetLastPrice(nil)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.LastPrice)
	assert.EqualValues(t, "234.22", resp.LastPrice.Usd)
}