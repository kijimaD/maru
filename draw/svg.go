package draw

import (
	"maru/config"
	"github.com/ajstarks/svgo"
	"os"
	"maru/utils"
)

type Draw struct {
	config config.Config
}

func New(config config.Config) Draw {
	return Draw {
		config: config,
	}
}

func (c Draw) Build() {
	out, err := os.Create("test.svg")
	utils.CheckErr(err)
	s := svg.New(out)
	s.Start(16, 16)
	s.Circle(8, 8, 8, "fill:#3ab60b;")
	s.End()
}
