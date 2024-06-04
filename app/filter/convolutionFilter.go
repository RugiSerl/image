package filter

import (
	"github.com/RugiSerl/image/app/image"
)

type weightMatrix = [][]float64

func getCoeffSum(matrix weightMatrix) float64 {
	s := 0.
	for _, column := range matrix {
		for _, coeff := range column {
			s += coeff
		}
	}
	return s
}

func convolutionFilter(img image.Image, weights weightMatrix) {
	imgWidth := len(img)
	imgHeight := len(img[0])

	// for each pixel
	for x, column := range img {
		for y := range column {

			// Use the convolution matrix
			s := image.Color{} // initialized to 0s
			for i := x - min(len(weights)/2-1, x); i < x+min(len(weights)/2-1, imgWidth-x); i++ {
				for j := y - min(len(weights)/2-1, y); j < y+min(len(weights)/2-1, imgHeight-y); j++ {
					s.AddColors(img[i][j])
				}
			}
			s.Scale(1. / getCoeffSum(weights))

			img[x][y] = s
		}
	}
}
