package math

type Polynomial []complex128 // Representation of polynomials with coefficients

func NewPolynomialFromReal(coefs []float64) Polynomial {
	return make(Polynomial, len(coefs))
}

// Returns the odd part of a polynomial
func (p Polynomial) Odd() Polynomial {
	r := make(Polynomial, len(p)/2)
	for i := 0; i < len(p)/2; i++ {
		r[i] = p[2*i+1]
	}
	return r
}

// Returns the even part of a polynomial
func (p Polynomial) Even() Polynomial {
	r := make(Polynomial, len(p)/2)
	for i := 0; i < len(p)/2; i++ {
		r[i] = p[2*i]
	}
	return r
}

// Evaluate the polynomial at a certain point
func (p Polynomial) Evaluate(z complex128) complex128 {
	s := 0 + 0i
	for i := 0; i < len(p); i++ {
		s = s + p[i]*Pow(z, i)
	}
	return s
}

// Fast exponentiation algorithm in O(log(n))
func Pow(z complex128, exponent int) complex128 {
	if exponent == 0 {
		return 1
	} else if exponent == 1 {
		return z
	} else if exponent%2 == 0 {
		r := Pow(z, exponent/2)
		return r * r
	} else if exponent%2 == 1 {
		r := Pow(z, exponent/2)
		return z * r * r
	} else {
		return 1
	}
}
