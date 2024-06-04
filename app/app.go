package app

import (
	"github.com/RugiSerl/image/app/filter"
	"github.com/RugiSerl/image/app/image"
)

func Run() {
	img, err := image.LoadImage("assets/sample.png")
	if err != nil {
		panic(err)
	}
	img = filter.Monochrome(img)
	img.SaveImage("assets/output.png")
}
