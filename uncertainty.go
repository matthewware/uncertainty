package uncertainty

import "math"

// UFloat represents a float value with associated uncertainty.
type UFloat struct {
	Value       float64 // Value of the float.
	Uncertainty float64 // Uncertainty of the float.
}

// Add returns the sum of two UFloat values.
func (uf1 UFloat) Add(uf2 UFloat) UFloat {
	return UFloat{
		Value:       uf1.Value + uf2.Value,
		Uncertainty: math.Sqrt(math.Pow(uf1.Uncertainty, 2) + math.Pow(uf2.Uncertainty, 2)),
	}
}

// Subtract returns the difference between two UFloat values.
func (uf1 UFloat) Subtract(uf2 UFloat) UFloat {
	return UFloat{
		Value:       uf1.Value - uf2.Value,
		Uncertainty: math.Sqrt(math.Pow(uf1.Uncertainty, 2) + math.Pow(uf2.Uncertainty, 2)),
	}
}

// Multiply returns the product of a UFloat value and a float64.
func (uf UFloat) MultiplyByFloat(f float64) UFloat {
	value := f * uf.Value
	uncertainty := math.Abs(f) * uf.Uncertainty
	return UFloat{
		Value:       value,
		Uncertainty: uncertainty,
	}
}

// Multiply returns the product of two UFloat values.
func (uf1 UFloat) Multiply(uf2 UFloat) UFloat {
	value := uf1.Value * uf2.Value
	uncertainty := math.Sqrt(math.Pow(uf1.Value*uf2.Uncertainty, 2) + math.Pow(uf2.Value*uf1.Uncertainty, 2))
	return UFloat{
		Value:       value,
		Uncertainty: uncertainty,
	}
}

// Divide returns the division of two UFloat values.
func (uf1 UFloat) Divide(uf2 UFloat) UFloat {
	value := uf1.Value / uf2.Value
	uncertainty := math.Sqrt(math.Pow(uf1.Uncertainty/uf2.Value, 2) + math.Pow(uf2.Uncertainty*uf1.Value/(math.Pow(uf2.Value, 2)), 2))
	return UFloat{
		Value:       value,
		Uncertainty: uncertainty,
	}
}
