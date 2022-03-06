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

	"github.com/sirupsen/logrus"
)

// AutoMigrate Create tables if not exists
func (s *Store) AutoMigrate(database string) error {

	if err := s.DBClient.Ping(context.Background()); err != nil {
		return err
	}

	// sql := ""
	// if database == "postgresql" {
	// 	file, err := ioutil.ReadFile("sql/postgresql.sql")
	// 	if err != nil {
	// 		return err
	// 	}
	// 	sql = string(file)
	// } else if database == "mysql" {
	// 	file, err := ioutil.ReadFile("sql/mysql.sql")
	// 	if err != nil {
	// 		return err
	// 	}
	// 	sql = string(file)
	// } else if database == "cockroachdb" {
	// 	file, err := ioutil.ReadFile("sql/cockroachdb.sql")
	// 	if err != nil {
	// 		return err
	// 	}
	// 	sql = string(file)
	// } else if database == "mongo" {
	// 	return nil
	// }

	// Create new table
	// builder := s.DBClient.SQL()
	// _, err := builder.Exec(sql)
	// if err != nil {
	// 	return err
	// }

	logrus.Info("Auto migrate completed")

	return nil
}
