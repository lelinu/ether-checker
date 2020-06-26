package services

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/lelinu/api_utils/log/lzap"
	"github.com/lelinu/api_utils/utils/error_utils"
	"github.com/lelinu/ether-checker/src/domain/account"
	"github.com/lelinu/ether-checker/src/domain/alert"
	"math/big"
	"strings"
)

type IAlertService interface {
	CheckBalance(req *alert.CheckBalanceRequest) *error_utils.ApiError
}

type alertService struct {
	logger         lzap.IService
	accountService IAccountService
	awsTopicArn    string
}

func NewAlertService(logger lzap.IService, accountService IAccountService, awsTopicArn string) IAlertService {
	return &alertService{logger: logger, accountService: accountService, awsTopicArn: awsTopicArn}
}

func (a *alertService) CheckBalance(req *alert.CheckBalanceRequest) *error_utils.ApiError {

	// validate request
	if err := req.Validate(); err != nil {
		return err
	}

	// get balances
	accounts, err := a.accountService.GetBalance(&account.GetBalanceRequest{Address: req.Address})
	if err != nil {
		return err
	}

	// build msg
	var msg strings.Builder

	// range into accounts
	for _, acc := range accounts.Accounts {

		value, _ := new(big.Float).Sub(acc.UsdValue, req.MinUsdValue).Float32()
		if value < 0 {
			msg.WriteString(fmt.Sprintf("Address: %v; Balance (Approx): $%v", acc.Address, acc.UsdValue))
		}
	}

	message := msg.String()
	// do not send an alert
	if len(message) == 0 {
		return nil
	}

	// create aws session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)
	_, awsErr := svc.Publish(&sns.PublishInput{
		Message:  &message,
		TopicArn: &a.awsTopicArn,
	})
	if awsErr != nil {
		return error_utils.NewInternalServerError(awsErr.Error())
	}

	return nil
}
