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
	"context"
	"fmt"

	"github.com/relayer/relayer/pkg/db"
	"github.com/relayer/relayer/pkg/db/adapter/postgresql"
)

// Store represents a store in the system
type Store struct {
	DBClient db.Adapter
}

// NewStore creates a new store
func NewStore(ctx context.Context, adapter string, connectionURL string) (*Store, error) {

	// Create new store instance
	store := &Store{}

	if adapter == "postgresql" {
		postgresqlAdapter, err := postgresql.NewPostgreSQLAdapter(connectionURL)
		if err != nil {
			return nil, err
		}
		store.DBClient = postgresqlAdapter
	} else {
		return nil, fmt.Errorf("unknown adapter: %s", adapter)
	}

	err := store.DBClient.Open(ctx)
	if err != nil {
		return nil, err
	}

	return store, nil
}

// Close closes the store
func (s *Store) Close() {
	s.DBClient.Close()
}
