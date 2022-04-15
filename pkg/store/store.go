package store

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	db *mongo.Client
}

func NewStore(ctx context.Context, mongoURI string) (*Store, error) {
	// connect to mongodb
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	// ping mongodb to check if it's up
	err = db.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}

func (s *Store) Close(ctx context.Context) error {
	return s.db.Disconnect(ctx)
}
