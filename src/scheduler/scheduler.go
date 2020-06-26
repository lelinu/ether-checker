package scheduler

import (
	"errors"
	"github.com/lelinu/api_utils/log/lzap"
	"github.com/lelinu/ether-checker/src/domain/alert"
	"github.com/lelinu/ether-checker/src/services"
	"github.com/lelinu/ether-checker/src/settings"
)

type IScheduler interface {
	RunCheckBalance()
}

type scheduler struct {
	logger       lzap.IService
	alertService services.IAlertService
	appSettings  settings.IAppSettings
}

func NewScheduler(logger lzap.IService, alertService services.IAlertService, appSettings settings.IAppSettings) IScheduler {
	return &scheduler{logger: logger, alertService: alertService, appSettings: appSettings}
}

func (s *scheduler) RunCheckBalance() {

	err := s.alertService.CheckBalance(&alert.CheckBalanceRequest{
		Address:     s.appSettings.GetEtherAddress(),
		MinUsdValue: s.appSettings.GetMinUsdValue(),
	})

	if err != nil {
		s.logger.Error("RunCheckBalance - An error had occurred.", errors.New(err.ErrorMessage))
	}
}
