package uncertainty

import (
	"math"
	"testing"
)

func TestUFloat_Add(t *testing.T) {
	uf1 := UFloat{Value: 10, Uncertainty: 1}
	uf2 := UFloat{Value: 5, Uncertainty: 0.5}
	expected := UFloat{Value: 15, Uncertainty: math.Sqrt(math.Pow(1, 2) + math.Pow(0.5, 2))}
	result := uf1.Add(uf2)

	if result != expected {
		t.Errorf("Addition test failed: expected %v, got %v", expected, result)
	}
}

func TestUFloat_Subtract(t *testing.T) {
	uf1 := UFloat{Value: 10, Uncertainty: 1}
	uf2 := UFloat{Value: 5, Uncertainty: 0.5}
	expected := UFloat{Value: 5, Uncertainty: math.Sqrt(math.Pow(1, 2) + math.Pow(0.5, 2))}
	result := uf1.Subtract(uf2)

	if result != expected {
		t.Errorf("Subtraction test failed: expected %v, got %v", expected, result)
	}
}

func TestScalar_Multiply(t *testing.T) {
	f1 := 2.00
	uf2 := UFloat{Value: 3, Uncertainty: 0.2}
	expected := UFloat{Value: 6, Uncertainty: 2 * 0.2}
	result := ScalarMultiply(f1, uf2)

	if result != expected {
		t.Errorf("Scalar multiplication test failed: expected %v, got %v", expected, result)
	}
}

func TestUFloat_Multiply(t *testing.T) {
	uf1 := UFloat{Value: 2, Uncertainty: 0.1}
	uf2 := UFloat{Value: 3, Uncertainty: 0.2}
	expected := UFloat{Value: 6, Uncertainty: math.Sqrt(math.Pow(2*0.2, 2) + math.Pow(3*0.1, 2))}
	result := uf1.Multiply(uf2)

	if result != expected {
		t.Errorf("Multiplication test failed: expected %v, got %v", expected, result)
	}
}

func TestUFloat_Divide(t *testing.T) {
	uf1 := UFloat{Value: 10, Uncertainty: 0.5}
	uf2 := UFloat{Value: 2, Uncertainty: 0.1}
	expected := UFloat{Value: 5, Uncertainty: math.Sqrt(math.Pow(0.5/2, 2) + math.Pow(0.1*10/(math.Pow(2, 2)), 2))}
	result := uf1.Divide(uf2)

	if result != expected {
		t.Errorf("Division test failed: expected %v, got %v", expected, result)
	}
}

// func ExampleUncertainty() {
//
// 	func foo(f1 UFloat, f2 UFloat) ( res UFloat) {
// 		return f1.Multiply(f2) // f1 * f2
// 	}
//
// 	func bar(f1 UFloat, f2 UFloat) (res UFloat) {
// 		return f1.Divide(f2).ScalarMulitply(4) // 4 * (f1 / f2)
// 	}
//
// 	// UFloat type is {value, standard deviation}
// 	a := UFloat(1.2, 0.3)
// 	b := UFloat(4.5, 0.6)
// 	fmt.Println(foo(a, b))
// 	fmt.Println(bar(a, b))
// }
