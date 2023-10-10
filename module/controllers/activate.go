package controllers

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/autobar-dev/services/module/repositories"
	"github.com/autobar-dev/services/module/types"
	"github.com/autobar-dev/services/module/utils"
	"github.com/google/uuid"
)

func ActivateController(
	app_context *types.AppContext,
	user_id string,
	serial_number string,
	otk string,
) error {
	rr := app_context.Repositories.Realtime
	sr := app_context.Repositories.State
	ur := app_context.Repositories.User
	wr := app_context.Repositories.Wallet
	cur := app_context.Repositories.Currency

	module, err := GetModuleController(app_context, serial_number)
	if err != nil {
		return err
	}

	if !module.Enabled {
		return errors.New("module not enabled")
	}

	module_otk, err := sr.GetOtkForModule(serial_number)
	if err != nil {
		return err
	}

	if *module_otk != otk {
		return errors.New("invalid otk")
	}

	user, err := ur.GetUserById(user_id)
	if err != nil {
		return err
	}

	user_wallet, err := wr.GetWallet(user_id)
	if err != nil {
		return err
	}

	currency, err := cur.GetCurrencyByCode(user_wallet.CurrencyCode)
	if err != nil {
		return err
	}

	if !currency.Enabled {
		return errors.New("currency not enabled")
	}

	var price_per_litre int
	if ppl, ok := module.Prices[currency.Code]; ok {
		price_per_litre = ppl
	} else { // if module has no specified price for this currency, get exchange rate
		var default_currency_code string
		var ppl_in_default_currency int

		for cc, ppl_c := range module.Prices {
			default_currency_code = cc
			ppl_in_default_currency = ppl_c
			break
		}

		if default_currency_code == "" || ppl_in_default_currency == 0 {
			return errors.New("prices for module set up incorrectly")
		}

		rate, err := cur.GetRate(default_currency_code, currency.Code)
		if err != nil {
			fmt.Printf("IMPORTANT: error getting exchange rate: %s\n", err)
			return errors.New("could not get exchange rate")
		}

		price_per_litre = int(math.Ceil(float64(ppl_in_default_currency) * rate.Rate))
	}

	args := &types.ActivateCommandArgs{
		UserInfo: types.ActivateCommandArgsUserInfo{
			FirstName: user.FirstName,
			Locale:    user.Locale,
			Wallet: types.ActivateCommandArgsUserInfoWallet{
				Balance: user_wallet.Balance,
				Currency: types.ActivateCommandArgsUserInfoWalletCurrency{
					Code:             currency.Code,
					Symbol:           currency.Symbol,
					MinorUnitDivisor: currency.MinorUnitDivisor,
				},
			},
		},
		PriceInfo: types.ActivateCommandArgsPriceInfo{
			PricePerLitre: price_per_litre,
		},
	}

	args_map, err := utils.StructToJsonMap(args)
	if err != nil {
		return err
	}

	err = rr.SendCommand(
		serial_number,
		repositories.ModuleServiceRealtimeClientType,
		types.ActivateCommandName,
		args_map,
	)
	if err != nil {
		return err
	}

	as_id := uuid.New().String()

	as := &types.ActivationSession{
		Id:                as_id,
		UserId:            user_id,
		SerialNumber:      serial_number,
		ProductId:         *module.ProductId,
		Price:             0,
		AmountMillilitres: 0,
		CreatedAt:         time.Now().UTC(),
		UpdatedAt:         time.Now().UTC(),
	}

	ras := utils.ActivationSessionToRedisActivationSession(*as)

	err = sr.SetActivationSession(as_id, ras)
	if err != nil {
		fmt.Printf("IMPORTANT: error setting activation session: %s\n", err.Error())
		return err
	}

	err = sr.SetActivationSessionIdForModule(serial_number, as_id)
	if err != nil {
		fmt.Printf("IMPORTANT: error setting activation session id for module: %s\n", err.Error())
		return err
	}

	err = sr.SetActivationSessionIdForUser(user_id, as_id)
	if err != nil {
		fmt.Printf("IMPORTANT: error setting activation session id for user: %s\n", err.Error())
		return err
	}

	return nil
}
