package rnums

import (
	"testing"
)

// Test the new system memory allocation function of NewSystem
func TestNewSystem(testCase *testing.T) {
	testCase.Log("To test that the RNS system is loaded with the correct values")

	numOfMods := 4
	redundant := int64(23)
	extension := int64(10)

	system := NewSystem(numOfMods, redundant, extension)
	if len(system.mods) != numOfMods+2 {
		testCase.Errorf("RNS Error: System not assigned the correct value")
	}
	if system.M < 0 {
		testCase.Errorf("RNS Error: System not assigned the correct value for M")
	}
}

// Test the new system memory allocation function of NewSystemFromMods
func TestNewSystemFromMods(testCase *testing.T) {
	testCase.Log("To test that the RNS system is loaded with the correct values")

	mods := []int64{3, 11, 13, 17}
	redundant := int64(23)
	extension := int64(10)

	system := NewSystemFromMods(mods, redundant, extension)
	if len(system.mods) != len(mods)+2 {
		testCase.Errorf("RNS Error: System not assigned the correct value")
	}
	if system.M < 0 {
		testCase.Errorf("RNS Error: System not assigned the correct value for M")
	}
}
