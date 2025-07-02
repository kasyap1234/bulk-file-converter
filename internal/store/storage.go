package store

import (
	"context"
	"io"
)

type Storage interface {
	Upload(ctx context.Context,bucket , key string, data io.Reader,contentType string)()
}