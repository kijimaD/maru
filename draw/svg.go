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
		path := fmt.Sprintf("images/%v.svg", k)
		color := fmt.Sprintf("%v", v.Color)
		d.image(path, color)
	}
}

func (d Draw) clean() {
	if _, err := os.Stat("images"); os.IsNotExist(err) {
		err = os.RemoveAll("images")
		utils.CheckErr(err)

		err = os.Mkdir("images", os.ModePerm)
	}
}

func (d Draw) image(filePath string, color string) {
	out, err := os.Create(filePath)
	utils.CheckErr(err)
	s := svg.New(out)
	s.Start(16, 16)
	s.Circle(8, 8, 8, fmt.Sprintf("fill:%v;", color))
	s.End()
}
