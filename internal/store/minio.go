package store

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioStorage struct {
	client *minio.Client
}

func NewStorage(ctx context.Context, endpoint, accessKeyID, secretAccessKey string, useSSL bool) (*MinioStorage, error) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatal(err)
	}
	return &MinioStorage{client: minioClient}, nil
}

func (m *MinioStorage) Upload(ctx context.Context, bucketName string, objectName string, reader io.Reader) error {
	_, err := m.client.PutObject(ctx, bucketName, objectName, reader, -1, minio.PutObjectOptions{})
	return err
}
func (m *MinioStorage) GetObject(ctx context.Context, bucketName string, objectName string, options minio.GetObjectOptions, destinationPath string) error {
	object, err := m.client.GetObject(ctx, bucketName, objectName, options)
	if err != nil {
		return fmt.Errorf("failed to get object: %w", err)
	}
	defer object.Close()

	file, err := os.Create(destinationPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, object)
	if err != nil {
		return fmt.Errorf("failed to copy object to file: %w", err)
	}

	return nil
}

func (m *MinioStorage) MakeBucket(ctx context.Context, bucketName string, options minio.MakeBucketOptions) error {
	err := m.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{
		Region: options.Region,
	})
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (m *MinioStorage) BucketExists(ctx context.Context, bucketName string) (bool, error) {
	found, err := m.client.BucketExists(ctx, bucketName)
	if err != nil {
		return false, err
	}
	return found, err
}
