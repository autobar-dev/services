package controllers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/autobar-dev/services/email/types"
)

func Send(ac *types.AppContext, from string, to string, subject string, plain_text string, html string) error {
	ep := ac.Providers.Email

	var retry_index int
	for retry_index = 0; retry_index < types.SendMaxRetryCount; retry_index++ {
		err := ep.Send(from, to, subject, plain_text, html)
		if err == nil {
			ac.Logger.Info(fmt.Sprintf("sent email successfully after retry %d", retry_index+1))
			return nil
		}

		ac.Logger.Error(fmt.Sprintf("failed to send email - retry %d, error", retry_index+1), err)
	}

	return errors.New("failed to send email after " + strconv.Itoa(types.SendMaxRetryCount) + " retries")
}
