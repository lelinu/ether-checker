package settings

import (
	"github.com/lelinu/api_utils/utils/env_utils"
	"math/big"
	"strconv"
)

type IAppSettings interface {
	GetMinUsdValue() *big.Float
	GetAwsTopicArn() string
	GetEtherScanBaseUrl() string
	GetEtherScanApiKey() string
	GetEtherAddress() string
}

type appSettings struct {
	MinUsdValue      *big.Float `json:"min_usd_value"`
	AwsTopicArn      string     `json:"aws_topic_arn"`
	EtherScanBaseUrl string     `json:"ether_scan_base_url"`
	EtherScanApiKey  string     `json:"ether_scan_api_key"`
	EtherAddress     string     `json:"ether_address"`
}

var (
	settings IAppSettings
)

func init() {

	value := env_utils.GetEnv("MIN_USD_VALUE", "60.0")
	f64MinUsdValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}

	settings = &appSettings{
		MinUsdValue:      big.NewFloat(f64MinUsdValue),
		AwsTopicArn:      env_utils.GetEnv("AWS_TOPIC_ARN", ""),
		EtherScanBaseUrl: env_utils.GetEnv("ETH_SCAN_BASE_URL", "https://api.etherscan.io/api"),
		EtherScanApiKey:  env_utils.GetEnv("ETH_SCAN_API_KEY", ""),
		EtherAddress:     env_utils.GetEnv("ETH_ADDRESS", ""),
	}
}

func New() IAppSettings {
	return settings
}

func (a *appSettings) GetMinUsdValue() *big.Float {
	return a.MinUsdValue
}

func (a *appSettings) GetAwsTopicArn() string {
	return a.AwsTopicArn
}

func (a *appSettings) GetEtherScanBaseUrl() string {
	return a.EtherScanBaseUrl
}

func (a *appSettings) GetEtherScanApiKey() string {
	return a.EtherScanApiKey
}

func (a *appSettings) GetEtherAddress() string {
	return a.EtherAddress
}
