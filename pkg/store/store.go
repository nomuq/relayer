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
	"io/ioutil"

	"github.com/sirupsen/logrus"
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

// Create tables if not exists
func (s *Store) AutoMigrate(database string) error {

	sql := ""
	if database == "postgresql" {
		file, err := ioutil.ReadFile("migrate/postgresql.sql")
		if err != nil {
			return err
		}
		sql = string(file)
	} else if database == "mysql" {
		file, err := ioutil.ReadFile("migrate/mysql.sql")
		if err != nil {
			return err
		}
		sql = string(file)
	} else if database == "cockroachdb" {
		file, err := ioutil.ReadFile("migrate/cockroachdb.sql")
		if err != nil {
			return err
		}
		sql = string(file)
	} else if database == "mongo" {
		return nil
	}

	// Create new table
	builder := s.Session.SQL()
	_, err := builder.Exec(sql)
	if err != nil {
		return err
	}

	logrus.Info("AutoMigrate: tables created")

	return nil
}
