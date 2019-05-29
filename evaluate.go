package rnums

// EvaluateRedundant returns errorRate and the wrong values
func EvaluateRedundant(mods []int64, redundant int64, extension int64) ([]int64, float64) {
	var missed = float64(0)
	var wrong []int64
	system := NewSystemFromMods(mods, redundant, extension)
	for num := int64(0); num < system.M; num++ {
		a := system.BinaryToRNS(float64(num))
		a = system.BaseExtension(a, extension)
		if a.fragments[len(system.mods)-1] != num%extension {
			wrong = append(wrong, num)
			missed++
		}
	}
	return wrong, (missed / float64(system.M))
}

// EvaluateModSetForRedundants is used to evaluate a modulus set for a given redundant modulus
func EvaluateModSetForRedundants(mods []int64, startRed int64, stopRed int64, extension int64) ([]float64, []float64) {
	var redundants []float64
	var errorRates []float64
	for redundant := startRed; redundant < stopRed; redundant++ {
		flag := false
		for _, mod := range mods {
			if int64(redundant)%mod == 0 || int64(redundant)%extension == 0 {
				flag = true
				break
			}
		}
		if flag {
			continue
		}

		_, errorRate := EvaluateRedundant(mods, int64(redundant), extension)
		if errorRate == 0 {
			for redundant < stopRed {
				errorRates = append(errorRates, 0)
				redundants = append(redundants, float64(redundant))
				redundant++
			}
			break
		}

		errorRates = append(errorRates, errorRate)
		redundants = append(redundants, float64(redundant))
	}
	return errorRates, redundants
}
