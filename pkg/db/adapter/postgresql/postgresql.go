package postgresql

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type PostgreSQLAdapter struct {
	connectionURL string
	connection    *pgx.Conn
}

func NewPostgreSQLAdapter(connectionURL string) (*PostgreSQLAdapter, error) {
	return &PostgreSQLAdapter{
		connectionURL: connectionURL,
	}, nil
}

func (adapter *PostgreSQLAdapter) Open(ctx context.Context) error {
	conn, err := pgx.Connect(ctx, adapter.connectionURL)
	if err != nil {
		return err
	}
	adapter.connection = conn
	return nil
}

func (adapter *PostgreSQLAdapter) Close() error {
	return adapter.connection.Close(context.Background())
}

func (adapter *PostgreSQLAdapter) IsClosed() bool {
	return adapter.connection.IsClosed()
}
