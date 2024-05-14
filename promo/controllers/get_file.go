package controllers

import (
	"github.com/autobar-dev/services/file/types"
	"github.com/autobar-dev/services/file/utils"
)

func GetFile(ac *types.AppContext, id string, download bool) (*types.File, error) {
	fr := *ac.Repositories.File
	sr := *ac.Repositories.S3

	pf, err := fr.Get(id)
	if err != nil {
		return nil, err
	}

	url, err := sr.GetFile(pf.Id, pf.Extension, download)
	if err != nil {
		return nil, err
	}

	file := utils.PostgresFileToFile(*pf, url)

	return file, nil
}
