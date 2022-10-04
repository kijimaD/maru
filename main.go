// read lang file code source: https://github.com/LeeReindeer/github-colors/blob/go/github-colors.go

package main

import (
	"maru/config"
	"maru/draw"
)

func main() {
	c := config.New()
	draw.New(c).Build()
}
