package utils

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/relayer/relayer/pkg/config"
)

func ColorTable() table.Writer {
	tw := table.NewWriter()
	// tw.Style().Color.Header = text.Colors{text.FgGreen}
	// tw.SetColumnConfigs(
	// 	[]table.ColumnConfig{
	// 		{Number: 1, Colors: text.Colors{text.FgYellow}},
	// 	},
	// )
	// tw.Style().Options.DrawBorder = false
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
