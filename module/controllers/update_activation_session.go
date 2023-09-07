package controllers

import (
	"time"

	"github.com/autobar-dev/services/module/types"
)

func UpdateActivationSessionController(
	app_context *types.AppContext,
	serial_number string,
	price int,
	amount_millilitres int,
) error {
	sr := app_context.Repositories.State

	as_id, err := sr.GetActivationSessionIdForModule(serial_number)
	if err != nil {
		return err
	}

	as, err := sr.GetActivationSession(*as_id)
	if err != nil {
		return err
	}

	as.Price = price
	as.AmountMillilitres = amount_millilitres
	as.UpdatedAt = time.Now().UTC()

	err = sr.SetActivationSession(*as_id, as)
	if err != nil {
		return err
	}

	return nil
}
