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

	file, err := fr.Get(id)
	if err != nil {
		return err
	}

	err = sr.DeleteFile(id, file.Extension)
	if err != nil {
		return err
	}

	err = fr.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
