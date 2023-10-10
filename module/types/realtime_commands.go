package types

import "github.com/autobar-dev/services/module/repositories"

const (
	RequestReportCommandName repositories.CommandName = "request-report"
	ActivateCommandName      repositories.CommandName = "activate"
	DeactivateCommandName    repositories.CommandName = "deactivate"
)

// Request report
type RequestReportCommandArgs struct {
	Queue string `json:"queue"`
}

// Activate
type ActivateCommandArgsUserInfoWalletCurrency struct {
	Code             string  `json:"code"`
	Symbol           *string `json:"symbol"`
	MinorUnitDivisor int     `json:"minor_unit_divisor"`
}

type ActivateCommandArgsUserInfoWallet struct {
	Balance  int                                       `json:"balance"`
	Currency ActivateCommandArgsUserInfoWalletCurrency `json:"currency"`
}

type ActivateCommandArgsUserInfo struct {
	FirstName string                            `json:"first_name"`
	Locale    string                            `json:"locale"`
	Wallet    ActivateCommandArgsUserInfoWallet `json:"wallet"`
}

type ActivateCommandArgsPriceInfo struct {
	PricePerLitre int `json:"price_per_litre"`
}

type ActivateCommandArgs struct {
	UserInfo  ActivateCommandArgsUserInfo  `json:"user_info"`
	PriceInfo ActivateCommandArgsPriceInfo `json:"price_info"`
}

// Deactivate
type DeactivateCommandArgs struct{}
