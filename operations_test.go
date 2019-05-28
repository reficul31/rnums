package rnums

import "testing"

// Test the base extension function
func TestBaseExtension(testCase *testing.T) {
	testCase.Log("To test that the base extension loads the base 10")

	redundant := int64(23)
	system := NewSystemFromMods([]int64{3, 7, 11, 13}, redundant)
	for num := int64(0); num < system.M; num++ {
		a := system.BinaryToRNS(float64(num))
		a = system.BaseExtension(a)

		if a.fragments[len(system.mods)-1] != num%10 {
			testCase.Errorf("RNS Error: Base correction incorrect for %d", num)
		}
	}
}

// Test the conversion from binary to RNS
func TestBinaryToRNS(testCase *testing.T) {
	testCase.Log("To test the binary to RNS conversion function")

	redundant := int64(23)
	system := NewSystem(8, redundant)
	rns := system.BinaryToRNS(-15.64)

	if rns.precision != -2 {
		testCase.Errorf("RNS Error: RNS value precision assigned is incorrect")
	}

	if rns.sign != -1 {
		testCase.Errorf("RNS Error: RNS sign bit not set")
	}

	var fragments = []int64{3, 2, 4, 0, 6, 0, 27, 14}
	for index, fragment := range fragments {
		if fragment != rns.fragments[index] {
			testCase.Errorf("RNS Error: Incorrect modulus values given back")
		}
	}
}

// Test the conversion from RNS to Binary
func TestRNSToBinary(testCase *testing.T) {
	testCase.Log("To test the conversion of a RNS number to binary")

	redundant := int64(23)
	number := -10.64
	system := NewSystem(8, redundant)
	rns := system.BinaryToRNS(number)
	result := system.RNSToBinary(rns)

	if number != result {
		testCase.Errorf("RNS Error: Returned number not of the correct binary form")
	}
}

// Test the multiplication of two RNS numbers
func TestMultiply(testCase *testing.T) {
	testCase.Log("To test the multiplication of two RNS numbers")

	redundant := int64(23)
	system := NewSystem(8, redundant)
	a := system.BinaryToRNS(3.73)
	b := system.BinaryToRNS(-10.6)
	c := system.BinaryToRNS(-39.538)

	mul := system.Multiply(a, b)
	if mul.precision != c.precision {
		testCase.Errorf("RNS Error: Precision of multiplication not accurate")
	}

	if mul.sign != c.sign {
		testCase.Errorf("RNS Error: RNS sign bit is not set")
	}

	for index := range mul.fragments {
		if mul.fragments[index] != c.fragments[index] {
			testCase.Errorf("RNS Error: Multiplication was unsuccessful")
		}
	}
}

// Test the addition function of two RNS numbers
func TestAdd(testCase *testing.T) {
	testCase.Log("To test addition of two RNS numbers")

	redundant := int64(23)
	system := NewSystem(8, redundant)
	a := system.BinaryToRNS(3.8)
	b := system.BinaryToRNS(10.2)
	c := system.BinaryToRNS(140)

	mul := system.Add(a, b)
	if mul.precision != -1 {
		testCase.Errorf("RNS Error: Precision of multiplication not accurate")
	}
	for index := range mul.fragments {
		if mul.fragments[index] != c.fragments[index] {
			testCase.Errorf("RNS Error: Multiplication was unsuccessful")
		}
	}
}

// Test the multiplicative inverse of a RNS number
func TestMultiplicativeInverse(testCase *testing.T) {
	testCase.Log("To test the multiplicative inverse of a number")

	redundant := int64(23)
	system := NewSystem(8, redundant)
	a := system.BinaryToRNS(-58.23)
	aInv := system.MultiplicativeInverse(a)

	if aInv.precision != a.precision {
		testCase.Errorf("RNS Error: Precision of Multiplicative Inverse not accurate")
	}

	product := system.Multiply(a, aInv)
	var fragments = []int64{1, 1, 1, 1, 1, 1, 1, 1}
	for index, fragment := range fragments {
		if fragment != product.fragments[index] {
			testCase.Errorf("RNS Error: Incorrect modulus values given back")
		}
	}
}

// Test the additive inverse of a RNS number
func TestAdditiveInverse(testCase *testing.T) {
	testCase.Log("To test the additive inverse of a number")

	redundant := int64(23)
	system := NewSystem(8, redundant)
	a := system.BinaryToRNS(10.64)
	aInv := system.AdditiveInverse(a)

	if aInv.precision != a.precision {
		testCase.Errorf("RNS Error: Precision of Multiplicative Inverse not accurate")
	}

	sum := system.Add(a, aInv)
	var fragments = []int64{0, 0, 0, 0, 0, 0, 0, 0}
	for index, fragment := range fragments {
		if fragment != sum.fragments[index] {
			testCase.Errorf("RNS Error: Incorrect modulus values given back")
		}
	}
}

// Test the division of two RNS numbers
func TestDivision(testCase *testing.T) {
	testCase.Log("To test the division of two numbers")

	redundant := int64(23)
	system := NewSystem(8, redundant)
	a := system.BinaryToRNS(4)
	b := system.BinaryToRNS(2)
	result := system.BinaryToRNS(2)

	div := system.Divide(a, b)
	if div.precision != 0 {
		testCase.Errorf("RNS Error: Precision of division not accurate")
	}
	for index := range result.fragments {
		if div.fragments[index] != result.fragments[index] {
			testCase.Errorf("RNS Error: Division was unsuccessful")
		}
	}
}

// Test the subtraction of two RNS numbers
func TestSubtract(testCase *testing.T) {
	testCase.Log("To test the subtraction of two numbers")

	redundant := int64(23)
	system := NewSystem(8, redundant)
	a := system.BinaryToRNS(20.643)
	b := system.BinaryToRNS(12.789)
	result := system.BinaryToRNS(7.85400)

	div := system.Subtract(a, b)
	if div.precision != result.precision {
		testCase.Errorf("RNS Error: Precision of division not accurate")
	}
	for index := range result.fragments {
		if div.fragments[index] != result.fragments[index] {
			testCase.Errorf("RNS Error: Division was unsuccessful")
		}
	}
}
