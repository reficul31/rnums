package rnums

import (
	"testing"
)

func TestGetErrorRate(testCase *testing.T) {
	testCase.Log("To test that the RNS system is loaded with the correct values")

	missed, errorRate := EvaluateRedundant([]int64{3, 7, 11}, int64(4))
	if errorRate != float64(0) && len(missed) != 0 {
		testCase.Errorf("RNS Error: Incorrect error rate computed for the system. Check:%d", missed[0])
	}
}
