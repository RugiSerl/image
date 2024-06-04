package filter

import (
	"github.com/RugiSerl/image/app/image"
)

func Monochrome(img image.Image) image.Image {
	for x, column := range img {
		for y, pixel := range column {
			avg := pixel.R/3 + pixel.G/3 + pixel.B/3
			img[x][y] = image.Color{
				R: avg,
				G: avg,
				B: avg,
				A: pixel.A,
			}
		}
	}
	return img
}
