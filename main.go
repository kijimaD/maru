// read lang file code source: https://github.com/LeeReindeer/github-colors/blob/go/github-colors.go

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
