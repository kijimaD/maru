package draw

import (
	"maru/config"
	"github.com/ajstarks/svgo"
	"os"
	"maru/utils"
	"fmt"
)

type Draw struct {
	config config.Config
}

func New(config config.Config) Draw {
	return Draw {
		config: config,
	}
}

func (d Draw) Build() {
	d.clean()

	for k, v := range d.config.Langs {
		d.dot(fmt.Sprintf("images/dot/%v.svg", k), v.Color)
		d.banner(fmt.Sprintf("images/banner/%v.svg", k), v.Color, k)
	}
	d.blankDot(fmt.Sprintf("images/dot/None.svg"), "#000000")
}

func (d Draw) clean() {
	err := os.RemoveAll("images")
	utils.CheckErr(err)

	err = os.Mkdir("images", os.ModePerm)
	utils.CheckErr(err)
	err = os.Mkdir("images/dot", os.ModePerm)
	utils.CheckErr(err)
	err = os.Mkdir("images/banner", os.ModePerm)
	utils.CheckErr(err)
}

func (d Draw) blankDot(filePath string, color string) {
	out, err := os.Create(filePath)
	utils.CheckErr(err)
	s := svg.New(out)
	s.Start(16, 16)
	s.Circle(8, 8, 7, fmt.Sprintf("fill:none;stroke:%v;", color))
	s.End()
}

func (d Draw) dot(filePath string, color string) {
	out, err := os.Create(filePath)
	utils.CheckErr(err)
	s := svg.New(out)
	s.Start(16, 16)
	s.Circle(8, 8, 8, fmt.Sprintf("fill:%v;", color))
	s.End()
}

func (d Draw) banner(filePath string, color string, lang string) {
	out, err := os.Create(filePath)
	utils.CheckErr(err)
	s := svg.New(out)
	s.Start(len(lang) * 8, 18)
	s.Roundrect(0, 0, len(lang) * 8, 18, 4, 4, fmt.Sprintf("fill:%v;", color))
	s.Gstyle("font-family:monospace;font-weight:bold;")
	s.Text(len(lang) * 8 / 2, 14, lang, fmt.Sprintf("fill:%v;font-size:12;text-anchor:middle", "#ffffff"))
	s.Gend()
	s.End()
}
