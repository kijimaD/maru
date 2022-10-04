package main

import (
	"maru/config"
	"maru/draw"
)

func main() {
	c := config.New()
	draw := draw.New(c)
	draw.Build()
	draw.WriteReadme()
}
