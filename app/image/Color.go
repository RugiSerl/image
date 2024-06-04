package image

import "image/color"

// normalized color values
type Color struct {
	R, G, B, A float64
}

type Image [][]Color

const N = 65535

func ColorInterfacetoColor(c color.Color) Color {
	r, g, b, a := c.RGBA()
	return Color{
		float64(r) / N,
		float64(g) / N,
		float64(b) / N,
		float64(a) / N,
	}
}

func (c Color) ToRGBA() color.RGBA {
	return color.RGBA{
		uint8(c.R * 255),
		uint8(c.G * 255),
		uint8(c.B * 255),
		uint8(c.A * 255),
	}

}
