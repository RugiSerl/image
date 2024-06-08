package math

import (
	"fmt"
	"math"
	"math/cmplx"
	"reflect"
)

func Gaussian(t, deviation float64) float64 {
	return 1 / math.Sqrt(2*math.Pi*deviation*deviation) * math.Exp(t*t/(2*deviation*deviation))
}

func Gaussian2D(x, y, deviation float64) float64 {
	//works, but no.
	//return Gaussian(x, deviation) * Gaussian(y, deviation)

	return 1 / (2 * math.Pi * deviation * deviation) * math.Exp((x*x+y*y)/(2*deviation*deviation))
}

func Clamp(value, min, max float64) float64 {
	if value > max {
		return max
	} else if value < min {
		return min
	} else {
		return value
	}
}

func Add(a, b interface{}) (interface{}, error) {
	value_a := reflect.ValueOf(a)
	value_b := reflect.ValueOf(b)

	if value_a.Kind() != value_b.Kind() {
		return nil, fmt.Errorf("Different kinds, can't add them.")
	}

	switch value_a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value_a.Int() + value_b.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return value_a.Uint() + value_b.Uint(), nil
	case reflect.Float32, reflect.Float64:
		return value_a.Float() + value_b.Float(), nil
	case reflect.String:
		return value_a.String() + value_b.String(), nil
	default:
		return nil, fmt.Errorf("Type does not support addition.")
	}
}

func Omega(t float64) complex128 {
	return cmplx.Exp(complex(0, 2*math.Pi*t))
}
