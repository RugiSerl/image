package image

func (c1 Color) Add(c2 Color) Color {
	return Color{
		c1.R + c2.R,
		c1.G + c2.G,
		c1.B + c2.B,
		c1.A + c2.A,
	}
}

func (c Color) Scale(scalar float64) Color {
	return Color{
		scalar * c.R,
		scalar * c.G,
		scalar * c.B,
		scalar * c.A,
	}
}
