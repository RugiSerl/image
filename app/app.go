package app

import (
	"github.com/RugiSerl/image/app/filter"
	"github.com/RugiSerl/image/app/image"
)

func Run() {
	img, err := image.LoadImage("assets/image.png")
	if err != nil {
		panic(err)
	}
	img = filter.ConvolutionFilter(img, filter.EdgeDetection())
	img.SaveImage("assets/output.png")
}
