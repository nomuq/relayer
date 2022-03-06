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
	"net"
	"os"
	"time"

	"github.com/relayer/relayer/pkg/config"
	"github.com/relayer/relayer/pkg/interceptor"
	"github.com/relayer/relayer/pkg/proto"
	"github.com/relayer/relayer/pkg/relayer"
	"github.com/relayer/relayer/pkg/store"
	"github.com/relayer/relayer/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	config := config.NewRelayerConfig()
	app := &cli.App{
		Name:        "relayer-server",
		Usage:       `High performance Instant messaging server.`,
		Description: `relayer-server is a high performance instant messaging server.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:      "config",
				Aliases:   []string{"c"},
				Usage:     "Load configuration from `FILE`",
				TakesFile: true,
			},
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
			&cli.IntFlag{
				Name:        "port",
				Usage:       "Port to listen on",
				EnvVars:     []string{"RELAYER_PORT"},
				Value:       1203,
				Destination: &config.Port,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "version",
				Usage: `Print the version number of relayer-server`,
				Action: func(c *cli.Context) error {
					utils.PrintVersion(version, commit, date)
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {

			// Print Relayer Logo and Version Info to the console
			fmt.Println(utils.RelayerLogo())

			// Load configuration from file
			err := config.Load(c.String("config"))
			if err != nil {
				return err
			}

			// If API key is not provided, generate a random one.
			if config.APIKey == "" {
				config.APIKey = utils.GenerateRandomString(15)
			}

			// If API secret is not provided, generate a random one.
			if config.APISecret == "" {
				config.APISecret = utils.GenerateRandomString(52)
			}

			// write the config to a file
			err = config.Write()
			if err != nil {
				return err
			}

			// Log the config.
			config.Print()

			// Initialize the database.
			store, err := store.NewStore(config.Database, config.DBConnectionURL)
			if err != nil {
				return err
			}
			defer store.Close()

			// Auto migrate the database. (Create tables if not exists)
			err = store.AutoMigrate(config.Database)
			if err != nil {
				return err
			}

			// Initialize interceptors.
			interceptor := interceptor.NewInterceptor(config)

			// Create new gRPC server
			server := grpc.NewServer(
				grpc.UnaryInterceptor(interceptor.UnaryInterceptor),
				grpc.StreamInterceptor(interceptor.StreamInterceptor),
			)

			// Register the relayer services with the gRPC server.
			relayer := relayer.NewRelayerServer(store)
			proto.RegisterRelayerServer(server, relayer)

			// listen on the port
			port := fmt.Sprintf("0.0.0.0:%d", config.Port)
			listener, err := net.Listen("tcp", port)
			if err != nil {
				return err
			}

			// start the server
			logrus.Infof("Starting server on port %d", config.Port)
			err = server.Serve(listener)
			if err != nil {
				return err
			}

			return nil
		},
	}

	// Run the app.
	if err := app.Run(os.Args); err != nil {
		// Log the error and exit.
		logrus.Errorln(err)
	}
}
