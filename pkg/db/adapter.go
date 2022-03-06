package db

import "context"

// generic database adapter interface for databases like postgresql, mysql, etc.
type Adapter interface {
	Open(ctx context.Context) error
	Close() error
	IsClosed() bool
}
