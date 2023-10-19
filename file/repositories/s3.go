package repositories

import (
	"context"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
)

type S3Repository struct {
	client      *minio.Client
	bucket_name string
}

func NewS3Repository(client *minio.Client, bucket_name string) *S3Repository {
	return &S3Repository{client, bucket_name}
}

func (sr *S3Repository) UploadFile(id string, extension string, file_header *multipart.FileHeader) error {
	file_size := file_header.Size
	file, err := file_header.Open()

	ctx := context.Background()
	_, err = sr.client.PutObject(
		ctx,
		sr.bucket_name,
		id+"."+extension,
		file,
		file_size,
		minio.PutObjectOptions{},
	)
	if err != nil {
		return err
	}

	return nil
}

func (sr *S3Repository) GetFile(id string, extension string, download bool) (read_url string, err error) {
	req_params := url.Values{}

	if download {
		req_params.Set("response-content-disposition", "attachment; filename=\""+id+"."+extension+"\"")
	} else {
		req_params.Set("response-content-disposition", "inline")
	}

	ctx := context.Background()
	url_object, err := sr.client.PresignedGetObject(
		ctx,
		sr.bucket_name,
		id+"."+extension,
		604800*time.Second,
		req_params,
	)
	if err != nil {
		return "", err
	}

	return url_object.String(), nil
}

func (sr *S3Repository) DeleteFile(id string, extension string) error {
	ctx := context.Background()
	return sr.client.RemoveObject(
		ctx,
		sr.bucket_name,
		id+"."+extension,
		minio.RemoveObjectOptions{},
	)
}
