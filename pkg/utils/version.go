package utils

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func PrintVersion(version string, commit string, date string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{
		"Version",
		"Commit",
		"Date",
	})
	t.AppendRow(table.Row{
		version,
		commit,
		date,
	})
	t.Render()
}
