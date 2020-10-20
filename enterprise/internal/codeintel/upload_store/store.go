package uploadstore

import (
	"context"
	"io"
)

// TODO - document
type Store interface {
	// TODO - document
	Init(ctx context.Context) error

	// TODO - document
	Get(ctx context.Context, key string, skipBytes int64) (io.ReadCloser, error)

	// TODO - document
	Upload(ctx context.Context, key string, r io.Reader) error
}

//
// TODO - new based on envvars
//
