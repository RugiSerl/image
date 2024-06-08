package math

import (
	"math"
)

type TimeDomainData = []complex128
type FrequencyDomainData = []complex128

//---------------------------------------------------------------------------------
// Mapping functions

// Annoying practical function because most of audio format samples are of type int
func MapIntArrayToTimeDomainData(samples []int) TimeDomainData {
	r := make(TimeDomainData, len(samples))
	for i := 0; i < len(samples); i++ {
		r[i] = complex(float64(samples[i]), 0)
	}
	return r
}

// Annoying practical function because most of audio format samples are of type int
func MapTimeDomainDataToIntArray(samples TimeDomainData) []int {
	r := make([]int, len(samples))
	for i := 0; i < len(samples); i++ {
		r[i] = int(real(samples[i]))
	}
	return r
}

// map fourier coefficients to time
func MapCoefsToTimeDomainData(coefficients FrequencyDomainData) TimeDomainData {
	r := make(TimeDomainData, len(coefficients))
	for i := 0; i < len(coefficients); i++ {
		r[i] = complex((1/float64(len(coefficients)))*real(coefficients[i])*math.Cos(2*math.Pi*float64(i))+real(coefficients[i])*math.Sin(2*math.Pi*float64(i)), 0)
	}
	return r
}

func AddZeroPadding[T any](slice []T, size int) []T {
	if size > len(slice) {
		t := make([]T, size) // golang will set everything to zero

		copy(t, slice)
		return t
	}
	return slice

}

//---------------------------------------------------------------------------------
// Fourrier functions

// Naive DTF algorithm in O(n²) (nested loops).
// Take time domain samples and returns frequency domains values.
// The values in the frequency domain are in that form :
// {amount of cosine of frequency} + i{amount of sine of frequency}.
// So to get the magnitude of the signal of the frequency a + ib,
// We calculate √(a²+b²) (Acos(ωt) + Bsin(ωt) = |A+iB|cos(ωt + phi)).
// So phi ≡ arg(a+ib) ≡ artan(b/a) [pi]
func DFTAux(samples TimeDomainData, inverse bool) FrequencyDomainData {
	N := len(samples)
	fourrierCoefficients := make([]complex128, N)
	//loop trough all frequencies possibilities
	for f := 0; f < N; f++ {

		//c[f] = ∑s[n]exp(-2πnf/N), n∈[0, N[
		for n := 0; n < N; n++ {
			var ω complex128
			if inverse {
				//inverse fourier transform is almost identical, just inverse the sign of ω and divide the sum by N
				ω = Omega(float64(f) * float64(n) / float64(N))
			} else {
				ω = Omega(-float64(f) * float64(n) / float64(N))
			}
			fourrierCoefficients[f] += samples[n] * ω
		}

	}

	return fourrierCoefficients
}

func DFT(samples TimeDomainData) FrequencyDomainData {
	return DFTAux(samples, true)
}

func InverseDFT(coefficients FrequencyDomainData) TimeDomainData {
	return MapCoefsToTimeDomainData(DFTAux(coefficients, false))
}

// Cooley Tuckey divide and conquer algorithm.
// Information and pseudo code here : https://fr.wikipedia.org/wiki/Transformation_de_Fourier_rapide#Pseudo-code
func FFTAux(samples Polynomial, ω complex128) FrequencyDomainData {
	N := len(samples)

	//constant polynomial case
	if N == 1 {
		return []complex128{samples[0]}
	} else {

		// calculate one time ω² to avoid unnecessary multiplications
		ω2 := Pow(ω, 2)

		// get recursive results from even and odd part of the polynomial
		evenResults := FFTAux(samples.Even(), ω2)
		oddResults := FFTAux(samples.Odd(), ω2)

		// merge back the result of the recursive results
		results := make([]complex128, N)
		ωk := 1 + 0i
		for k := 0; k < N/2; k++ {
			results[k] = evenResults[k] + ωk*oddResults[k]
			results[k+N/2] = evenResults[k] - ωk*oddResults[k]

			ωk *= ω // avoid to call Pow(ω, k)
		}

		return results
	}

}

func FFT(samples TimeDomainData) FrequencyDomainData {
	return FFTAux(samples, Omega(-1/float64(len(samples))))
}

func InverseFFT(coefficients FrequencyDomainData) TimeDomainData {
	return MapCoefsToTimeDomainData(FFTAux(coefficients, Omega(1/float64(len(coefficients)))))
}

// Double computation of the FFT in a matrix.
// Since the fourier transform is linear, we can just do 2 fourrier transform
func FFT2D(m ComplexMatrix) ComplexMatrix {
	rowsTransformed := InitMatrix[complex128](len(m), len(m[0]))
	for i := 0; i < len(m); i++ {
		rowsTransformed[i] = FFT(m[i])
	}
	return rowsTransformed

}
