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

package utils

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/relayer/relayer/pkg/config"
)

func ColorTable() table.Writer {
	tw := table.NewWriter()
	tw.SetOutputMirror(os.Stdout)
	tw.SetStyle(table.StyleLight)
	return tw
}

func LogConfig(conf config.RelayerConfig) {
	t := ColorTable()
	t.AppendHeader(table.Row{
		"JWT secret key",
		"API key",
		"API secret",
	})
	t.AppendRow(table.Row{
		conf.JWTSecret,
		conf.APIKey,
		conf.APISecret,
	})
	t.Render()
}
