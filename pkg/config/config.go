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

type RelayerConfig struct {
	APIKey          string `yaml:"api_key"`
	APISecret       string `yaml:"api_secret"`
	Database        string `yaml:"database"`
	DBConnectionURL string `yaml:"db_connection_url"`
	path            string
}

func NewRelayerConfig() *RelayerConfig {
	return &RelayerConfig{}
}

func (c *RelayerConfig) Load(path string) error {
	c.path = path
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
		fmt.Println(err)
		return err
	}

	return nil
}

func (c *RelayerConfig) LoadYAML() error {
	yfile, err := ioutil.ReadFile(c.path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yfile, &c)
	if err != nil {
		return err
	}
	return nil
}

func (c *RelayerConfig) Write() error {
	data, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(c.path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ColorTable() table.Writer {
	tw := table.NewWriter()
	tw.SetOutputMirror(os.Stdout)
	tw.SetStyle(table.StyleLight)
	return tw
}

func (c *RelayerConfig) Print() {
	t := ColorTable()
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
