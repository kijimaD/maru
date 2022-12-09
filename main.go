package main

import (
	"github.com/kijimaD/maru/config"
	"github.com/kijimaD/maru/draw"
)

func main() {
	c := config.New()
	draw := draw.New(c)
	draw.Build()
	draw.WriteReadme()
}
