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

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	cfg "github.com/relayer/relayer/pkg/config"
	str "github.com/relayer/relayer/pkg/store"
	"github.com/relayer/relayer/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	config := cfg.RelayerConfig{}
	app := &cli.App{
		Name:        "relayer-server",
		Usage:       `High performance Instant messaging server.`,
		Description: `relayer-server is a high performance instant messaging server.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "api-key",
				Usage:       "API key for relayer-server",
				EnvVars:     []string{"RELAYER_API_KEY"},
				Value:       "",
				Destination: &config.APIKey,
			},
			&cli.StringFlag{
				Name:        "api-secret",
				Usage:       "API secret for relayer-server",
				EnvVars:     []string{"RELAYER_API_SECRET"},
				Value:       "",
				Destination: &config.APISecret,
			},
			&cli.StringFlag{
				Name:        "database",
				Usage:       "Database type",
				EnvVars:     []string{"RELAYER_DATABASE"},
				Value:       "postgresql",
				Destination: &config.Database,
			},
			&cli.StringFlag{
				Name:        "db-connection-url",
				Usage:       "Database connection URL",
				EnvVars:     []string{"RELAYER_DB_CONNECTION_URL"},
				Value:       "postgres://postgres:postgres@localhost:5432/relayer?sslmode=disable",
				Destination: &config.DBConnectionURL,
			},
		},
		Commands: []*cli.Command{},
		Action: func(c *cli.Context) error {
			fmt.Println(utils.RelayerLogo())
			// If API key is not provided, generate a random one.
			if config.APIKey == "" {
				config.APIKey = utils.GenerateRandomString(15)
			}

			// If API secret is not provided, generate a random one.
			if config.APISecret == "" {
				config.APISecret = utils.GenerateRandomString(52)
			}

			// Log the config.
			utils.LogConfig(config)

			// Initialize the database.
			store, err := str.NewStore(config.Database, config.DBConnectionURL)
			if err != nil {
				logrus.Fatalf("failed to create database: %v", err)
			}
			defer store.Close()

			err = store.AutoMigrate(config.Database)
			if err != nil {
				logrus.Fatalf("failed to migrate database: %v", err)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Errorln(err)
	}
}
