package draw

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
	"os"
	"maru/utils"
)

func (d Draw) WriteReadme() {
	file, err := os.Create("README.md")
	utils.CheckErr(err)
	defer file.Close()

	d.header(file)
	d.colorTable(file)
}

func (d *Draw) header(w io.Writer) {
	str := fmt.Sprintf("# maru\n\n")
	str += fmt.Sprintf("\n")
	w.Write([]byte(str))
}

func (d Draw) colorTable(w io.Writer) {
	table := tablewriter.NewWriter(w)
	table.SetHeader(d.colorHeader())
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(false)

	for k, v := range d.config.Langs {
		table.Append(d.colorContent(k, v.Color))
	}
	table.Render()
}

func (d Draw) colorHeader() []string {
	return []string{
		"Lang",
		"Image",
	}
}

func (d Draw) colorContent(name string, color string) []string {
	return []string{
		name,
		fmt.Sprintf("![%v](https://raw.githubusercontent.com/kijimaD/maru/main/images/%v.svg)", name, name),
	}
}
