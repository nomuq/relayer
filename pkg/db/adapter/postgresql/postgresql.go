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

package postgresql

import (
	"context"
	"io/ioutil"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
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

func (adapter *PostgreSQLAdapter) AutoMigrate(ctx context.Context) error {

	file, err := ioutil.ReadFile("sql/postgresql.sql")
	if err != nil {
		return err
	}
	sql := string(file)

	_, err = adapter.connection.Exec(ctx, sql)
	if err != nil {
		return err
	}

	logrus.Info("Auto migrate completed")

	return nil
}
