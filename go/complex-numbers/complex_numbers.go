package complexnumbers

import "math"

// Number type here.
type Number struct {
	real      float64
	imaginary float64
}

func New(real, imaginary float64) Number {
	return Number{real: real, imaginary: imaginary}
}

func (n Number) Real() float64 {
	return n.real
}

func (n Number) Imaginary() float64 {
	return n.imaginary
}

func (n Number) Add(m Number) Number {
	return New(n.Real()+m.Real(), n.Imaginary()+m.Imaginary())
}

func (n Number) Subtract(m Number) Number {
	return New(n.Real()-m.Real(), n.Imaginary()-m.Imaginary())
}

func (n Number) Multiply(m Number) Number {
	return New(
		n.Real()*m.Real()-n.Imaginary()*m.Imaginary(),
		n.Imaginary()*m.Real()+n.Real()*m.Imaginary(),
	)
}

func (n Number) Times(factor float64) Number {
	return New(n.Real()*factor, n.Imaginary()*factor)
}

func (n Number) Divide(m Number) Number {
	d := m.Real()*m.Real() + m.Imaginary()*m.Imaginary()
	r := (n.Real()*m.Real() + n.Imaginary()*m.Imaginary()) / d
	i := (n.Imaginary()*m.Real() - n.Real()*m.Imaginary()) / d
	return New(r, i)
}

func (n Number) Conjugate() Number {
	return New(n.Real(), -n.Imaginary())
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.Real()*n.Real() + n.Imaginary()*n.Imaginary())
}

func (n Number) Exp() Number {
	return New(math.Cos(n.Imaginary()), math.Sin(n.Imaginary())).Times(math.Exp(n.Real()))
}
