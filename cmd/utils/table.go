package utils

import (
	"io"

	"github.com/olekukonko/tablewriter"
)

func NewTable(writer io.Writer, headerElements []string) *tablewriter.Table {
	table := tablewriter.NewWriter(writer)
	table.SetHeader(headerElements)
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)

	return table
}
