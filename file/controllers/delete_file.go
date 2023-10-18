package controllers

import (
	"github.com/autobar-dev/services/file/types"
)

func DeleteFile(
	ac *types.AppContext,
	id string,
) error {
	fr := ac.Repositories.File
	sr := ac.Repositories.S3

	err := sr.DeleteFile(id)
	if err != nil {
		return err
	}

	err = fr.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
