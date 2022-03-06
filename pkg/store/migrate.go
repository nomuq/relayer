package store

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

// Create tables if not exists
func (s *Store) AutoMigrate(database string) error {

	sql := ""
	if database == "postgresql" {
		file, err := ioutil.ReadFile("sql/postgresql.sql")
		if err != nil {
			return err
		}
		sql = string(file)
	} else if database == "mysql" {
		file, err := ioutil.ReadFile("sql/mysql.sql")
		if err != nil {
			return err
		}
		sql = string(file)
	} else if database == "cockroachdb" {
		file, err := ioutil.ReadFile("sql/cockroachdb.sql")
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

	logrus.Info("Auto migrate completed")

	return nil
}
