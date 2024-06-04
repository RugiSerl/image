package image

import (
	"image"
	"image/png"
	"os"
)

// Init bidimensionnal slice, allocating the needed space. No need to free thanks to garbage collector.
func InitSlice[T any](x, y int) [][]T {
	result := make([][]T, x)
	for i := range result {
		result[i] = make([]T, y)
	}
	return result
}

func (img Image) GetDimension() image.Point {
	return image.Point{
		len(img),
		len(img[0]),
	}
}

func LoadImage(filePath string) (Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	var imageMatrix Image = InitSlice[Color](width, height)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			imageMatrix[x][y] = ColorInterfacetoColor(img.At(x, y))
		}
	}

	return imageMatrix, nil
}

func (imageMatrix Image) SaveImage(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, imageMatrix.GetDimension()})

	width := imageMatrix.GetDimension().X
	height := imageMatrix.GetDimension().Y

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, imageMatrix[x][y].ToRGBA())
		}
	}

	// Encode to `PNG` with `DefaultCompression` level
	// then save to file
	err = png.Encode(f, img)
	if err != nil {
		return err
	}

	return nil
}
