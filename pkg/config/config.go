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

package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"gopkg.in/yaml.v3"
)

// RelayerConfig represents the relayer config
type RelayerConfig struct {
	// APIKey is the API key for the relayer
	APIKey string `yaml:"api_key"`
	// APISecret is the API secret for the relayer
	APISecret string `yaml:"api_secret"`
	// Database is the type of database to use
	Database string `yaml:"database"`
	// DatabaseHost is the connection string for the database
	DBConnectionURL string `yaml:"db_connection_url"`
	// Port is the port to listen on
	Port int `yaml:"port"`
	// ConfigPath is the path to the config file
	path string
}

// NewRelayerConfig returns a new relayer config
func NewRelayerConfig() *RelayerConfig {
	return &RelayerConfig{}
}

// Load loads the config from the given path
func (c *RelayerConfig) Load(path string) error {

	// set path for config file
	c.path = path

	// if config file doesn't exist, create it
	if c.path == "" {
		c.path = "./config.yaml"
	}

	// check if file exists if not create it
	if _, err := os.Stat(c.path); os.IsNotExist(err) {
		// create file at path
		_, err = os.Create(c.path)
		if err != nil {
			return err
		}
	}

	// load config file
	err := c.LoadYAML()
	if err != nil {
		return err
	}

	return nil
}

// LoadYAML loads and parses the config file
func (c *RelayerConfig) LoadYAML() error {
	// read file contents into byte array
	yfile, err := ioutil.ReadFile(c.path)
	if err != nil {
		return err
	}

	// unmarshal yaml
	err = yaml.Unmarshal(yfile, &c)
	if err != nil {
		return err
	}
	return nil
}

// Save saves the config to the given path
func (c *RelayerConfig) Write() error {
	// marshal config
	data, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}

	// write to file
	err = ioutil.WriteFile(c.path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Print prints the config to the console
func (c *RelayerConfig) Print() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{
		"API key",
		"API secret",
	})
	t.AppendRow(table.Row{
		c.APIKey,
		c.APISecret,
	})
	t.Render()

	// Add Extra new line after table
	fmt.Println()
}
