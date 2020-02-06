package print

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"

	"github.com/tobbbles/muzip/archive"
)

func Pretty(archives []*archive.Archive) {
	data := [][]string{}

	for _, a := range archives {
		for _, content := range a.Contents {
			data = append(data, []string{a.Attr.Name, content.Artist, fmt.Sprintf("%d - %s", content.Track, content.Title)})
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeader([]string{"Archive", "Artist", "# Title"})
	table.SetBorder(false)
	table.SetAutoMergeCells(true)
	table.SetTablePadding("\t")
	table.SetRowSeparator("-")
	table.SetRowLine(true)

	table.AppendBulk(data)
	table.Render()
}
