/*
 * (C) Copyright 2022 Satish Babariya (https://satishbabariya.com/) and others.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Contributors:
 *     satish babariya (satish.babariya@gmail.com)
 *
 */

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

// Store represents a store in the system
type Store struct {
	DBClient db.Session
}

// NewStore creates a new store
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
		store.DBClient = session
	} else if adapter == "mysql" {
		dsn, err := mysql.ParseURL(connectionURL)
		if err != nil {
			return nil, err
		}
		session, err := mysql.Open(dsn)
		if err != nil {
			return nil, err
		}
		store.DBClient = session
	} else if adapter == "cockroachdb" {
		url, err := cockroachdb.ParseURL(connectionURL)
		if err != nil {
			return nil, err
		}
		session, err := cockroachdb.Open(url)
		if err != nil {
			return nil, err
		}
		store.DBClient = session
	} else if adapter == "mongo" {
		url, err := mongo.ParseURL(connectionURL)
		if err != nil {
			return nil, err
		}
		session, err := mongo.Open(url)
		if err != nil {
			return nil, err
		}
		store.DBClient = session
	} else {
		return nil, fmt.Errorf("unknown adapter: %s", adapter)
	}

	db.LC().SetLevel(db.LogLevelError)

	return store, nil
}

// Close closes the store
func (s *Store) Close() {
	s.DBClient.Close()
}
