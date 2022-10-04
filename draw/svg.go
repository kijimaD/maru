package draw

import (
	"maru/config"
	// "github.com/ajstarks/svgo"
)

type Draw struct {
	config config.Config
}

func New(config config.Config) Draw {
	return Draw {
		config: config,
	}
}
