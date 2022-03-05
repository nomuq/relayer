package store

import (
	"fmt"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/cockroachdb"
	"github.com/upper/db/v4/adapter/mongo"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
)

// Create generic store adapter interface for database postgresql, mysql, etc.

type Store struct {
	Session db.Session
}

func NewStore(adapter string, connectionURL string) (*Store, error) {

	// Create new store instance
	store := &Store{}

	if adapter == "postgresql" {
		url, err := postgresql.ParseURL(connectionURL)
		if err != nil {
			return nil, err
		}
		session, err := postgresql.Open(url)
		if err != nil {
			return nil, err
		}
		store.Session = session
	} else if adapter == "mysql" {
		dsn, err := mysql.ParseURL(connectionURL)
		if err != nil {
			return nil, err
		}
		session, err := mysql.Open(dsn)
		if err != nil {
			return nil, err
		}
		store.Session = session
	} else if adapter == "cockroachdb" {
		url, err := cockroachdb.ParseURL(connectionURL)
		if err != nil {
			return nil, err
		}
		session, err := cockroachdb.Open(url)
		if err != nil {
			return nil, err
		}
		store.Session = session
	} else if adapter == "mongo" {
		url, err := mongo.ParseURL(connectionURL)
		if err != nil {
			return nil, err
		}
		session, err := mongo.Open(url)
		if err != nil {
			return nil, err
		}
		store.Session = session
	} else {
		return nil, fmt.Errorf("unknown adapter: %s", adapter)
	}

	return store, nil
}

func (s *Store) Close() {

	s.Session.Close()
}
