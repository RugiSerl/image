package math

type Matrix[T any] [][]T

// Init bidimensionnal slice, allocating the needed space. No need to free thanks to garbage collector.
func InitMatrix[T any](x, y int) [][]T {
	result := make([][]T, x)
	for i := range result {
		result[i] = make([]T, y)
	}
	return result
}

func (m Matrix[T]) GetWidth() int {
	return len(m)
}
func (m Matrix[T]) GetHeight() int {
	return len(m[0])
}

func (m Matrix[T]) Add(n Matrix[T]) Matrix[T] {
	if len(m) != len(n) || len(m[0]) != len(n[0]) {
		panic("invalid matrix sizes")
	}
	var temp interface{}
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			temp, _ = Add(m[i][j], n[i][j])
			m[i][j] = temp.(T) // done in two steps because of error "handling"

		}
	}
	return m
}

// Matrix with real coefficients
type RealMatrix Matrix[float64]

// Matrix with complex coefficients
type ComplexMatrix Matrix[complex128]

func (m RealMatrix) ToComplex() ComplexMatrix {
	result := InitMatrix[complex128](len(m), len(m[0]))
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			result[i][j] = complex(m[i][j], 0)
		}
	}
	return result
}

func (m RealMatrix) Scale(scalar float64) RealMatrix {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			m[i][j] *= scalar
		}
	}
	return m
}
