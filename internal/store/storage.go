package store

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
)

type Storage interface {
	Upload(ctx context.Context, bucket, key string, data io.Reader, contentType string) error
	GetObject(ctx context.Context, bucketName string, objectName string, options minio.GetObjectOptions, destinationPath string)
	MakeBucket(ctx context.Context, bucketName string, options minio.MakeBucketOptions)
	BucketExists(ctx context.Context, bucketName string) (bool, error)
}
