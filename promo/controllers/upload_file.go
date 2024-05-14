package controllers

import (
	"mime/multipart"

	"github.com/google/uuid"

	"github.com/autobar-dev/services/file/types"
	"github.com/autobar-dev/services/file/utils"
)

func UploadFile(
	ac *types.AppContext,
	file_header *multipart.FileHeader,
) (id string, err error) {
	fr := ac.Repositories.File
	sr := ac.Repositories.S3

	// fmt.Println("File size:", file_header.Size)
	// fmt.Println("File name:", file_header.Filename)
	//
	// mime_version := file_header.Header.Get("MIME-Version")
	// content_type := file_header.Header.Get("Content-Type")
	// content_transfer_encoding := file_header.Header.Get("Content-Transfer-Encoding")
	// content_id := file_header.Header.Get("Content-ID")
	// content_description := file_header.Header.Get("Content-Description")
	// content_disposition := file_header.Header.Get("Content-Disposition")
	//
	// fmt.Println("MIME-Version:", mime_version)
	// fmt.Println("Content-Type:", content_type)
	// fmt.Println("Content-Transfer-Encoding:", content_transfer_encoding)
	// fmt.Println("Content-ID:", content_id)
	// fmt.Println("Content-Description:", content_description)
	// fmt.Println("Content-Disposition:", content_disposition)

	uuid := uuid.NewString()
	file_extension := utils.FileExtensionFromFileName(file_header.Filename)

	err = sr.UploadFile(uuid, file_extension, file_header)
	if err != nil {
		return "", err
	}

	err = fr.Create(uuid, file_extension)
	if err != nil {
		return "", err
	}

	return uuid, nil
}
