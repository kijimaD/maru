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
		d.image(fmt.Sprintf("images/dot/%v.svg", k), v.Color)
		d.banner(fmt.Sprintf("images/banner/%v.svg", k), k, v.Color)
	}
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

func (d Draw) image(filePath string, color string) {
	out, err := os.Create(filePath)
	utils.CheckErr(err)
	s := svg.New(out)
	s.Start(16, 16)
	s.Circle(8, 8, 8, fmt.Sprintf("fill:%v;", color))
	s.End()
}

func (d Draw) banner(filePath string, lang string, color string) {
	out, err := os.Create(filePath)
	utils.CheckErr(err)
	s := svg.New(out)
	s.Start(len(lang) * 10, 20)
	s.Roundrect(0, 0, len(lang) * 10, 20, 4, 4, fmt.Sprintf("fill:%v;", color))
	s.Text(4, 14, lang, fmt.Sprintf("fill:%v;font-size:12", "#ffffff"))
	s.End()
}
