package math

import "math"

func Gaussian(t, deviation float64) float64 {
	return 1 / math.Sqrt(2*math.Pi*deviation*deviation) * math.Exp(t*t/(2*deviation*deviation))
}

func Gaussian2D(x, y, deviation float64) float64 {
	return Gaussian(x, deviation) * Gaussian(y, deviation)
}
