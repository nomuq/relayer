package postgresql

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgreSQLAdapter struct {
	connectionURL string
	connection    *pgxpool.Pool
}

func NewPostgreSQLAdapter(connectionURL string) (*PostgreSQLAdapter, error) {
	return &PostgreSQLAdapter{
		connectionURL: connectionURL,
	}, nil
}

func (adapter *PostgreSQLAdapter) Open(ctx context.Context) error {
	connection, err := pgxpool.Connect(ctx, adapter.connectionURL)
	if err != nil {
		return err
	}
	adapter.connection = connection
	return nil
}

func (adapter *PostgreSQLAdapter) Close() {
	adapter.connection.Close()
}

func (adapter *PostgreSQLAdapter) Ping(ctx context.Context) error {
	return adapter.connection.Ping(ctx)
}
