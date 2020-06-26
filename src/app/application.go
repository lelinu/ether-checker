package app

import (
	"fmt"
	"github.com/lelinu/api_utils/log/lzap"
	"github.com/lelinu/ether-checker/src/providers/ether_scan_provider"
	"github.com/lelinu/ether-checker/src/scheduler"
	"github.com/lelinu/ether-checker/src/services"
	"github.com/lelinu/ether-checker/src/settings"
)

func StartApplication(){

	fmt.Println("start running")

	// init settings
	appSettings := settings.New()

	// init provider
	etherScanProvider := ether_scan_provider.NewProvider(appSettings.GetEtherScanBaseUrl(), appSettings.GetEtherScanApiKey())

	// init logger
	logger := lzap.NewService("info", "")

	// init services
	accountService := services.NewAccountService(logger, etherScanProvider)
	alertService := services.NewAlertService(logger, accountService, appSettings.GetAwsTopicArn())

	// runner
	runner := scheduler.NewScheduler(logger, alertService, appSettings)
	runner.RunCheckBalance()

	fmt.Println("stopped running")
}
