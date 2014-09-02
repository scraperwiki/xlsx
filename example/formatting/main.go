package main

import (
	"os"

	"github.com/psmithuk/xlsx"
)

func main() {
	outputfile, err := os.Create("test.xlsx")
	if err != nil {
		panic(err)
	}

	ww := xlsx.NewWorkbookWriter(outputfile)

	c := []xlsx.Column{
		xlsx.Column{Name: "Col1", Width: 10},
		xlsx.Column{Name: "Col2", Width: 10},
		xlsx.Column{Name: "Col3", Width: 10},
	}

	sh := xlsx.NewSheetWithColumns(c)
	sh.Title = "MySheet"

	sw, err := ww.NewSheetWriter(&sh)
	if err != nil {
		panic(err)
	}

	r := sh.NewRow()
	r.Cells[0] = xlsx.Cell{
		Type:    xlsx.CellTypeInlineString,
		Value:   "Spanning Title",
		Colspan: 2,
	}
	r.Cells[2] = xlsx.Cell{
		Type:    xlsx.CellTypeInlineString,
		Value:   "Hello",
		Rowspan: 2,
	}
	err = sw.WriteRows([]xlsx.Row{r})
	if err != nil {
		panic(err)
	}

	r = sh.NewRow()
	r.Cells[0] = xlsx.Cell{
		Type:  xlsx.CellTypeNumber,
		Value: "5",
	}
	r.Cells[1] = xlsx.Cell{
		Type:  xlsx.CellTypeNumber,
		Value: "10",
	}
	err = sw.WriteRows([]xlsx.Row{r})
	if err != nil {
		panic(err)
	}

	err = ww.Close()
	if err != nil {
		panic(err)
	}

	defer outputfile.Close()
}
