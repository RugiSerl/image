package filter

import (
	"github.com/RugiSerl/image/app/image"
	"github.com/RugiSerl/image/app/math"
)

type weightMatrix = math.RealMatrix

func ConvolutionFilter(img image.Image, weights weightMatrix) image.Image {
	imgWidth := len(img)
	imgHeight := len(img[0])

	resultImg := math.InitMatrix[image.Color](len(img), len(img[0]))

	// for each pixel
	for x, column := range img {
		for y := range column {

			// Use the convolution matrix
			s := image.Color{} // initialized to 0s
			WeightSum := 0.    // sum of the weight used
			for i := 0; i < len(weights); i++ {
				for j := 0; j < len(weights); j++ {

					onImageX := i - len(weights)/2 - 1 + x
					onImageY := j - len(weights)/2 - 1 + y

					// Is it in the image
					if onImageX >= 0 && // left
						onImageX < imgWidth && // right
						onImageY >= 0 && // top
						onImageY < imgHeight { // bottom

						w := weights[i][j]
						s = s.Add(img[onImageX][onImageY].Scale(w))
						WeightSum += w

					}

				}
			}

			if WeightSum != 0 {
				s = s.Scale(1. / WeightSum)
			}

			resultImg[x][y] = s.Clamp()
		}
	}
	return resultImg
}

func GaussianBlurMatrix(blurSize int) weightMatrix {
	var mat weightMatrix = math.InitMatrix[float64](2*blurSize+1, 2*blurSize+1)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			mat[i][j] = math.Gaussian2D(float64(i-len(mat)/2), float64(j-len(mat)/2), float64(blurSize)*3)
		}
	}
	return mat
}

func EdgeDetection() weightMatrix {
	var mat weightMatrix = weightMatrix{
		{-1, -1, -1},
		{-1, 5, -1},
		{-1, -1, -1},
	}
	return mat
}
