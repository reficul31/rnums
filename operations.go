package rnums

import (
	"fmt"
	"math"
	"strconv"
)

// BaseExtension regenerates the base 10 for a RNS number
func (system System) BaseExtension(a *RNS, extension int64) *RNS {
	sum := []int64{0, 0}
	redundant := system.mods[len(system.mods)-2]
	for i := 0; i < len(system.mods)-2; i++ {
		sum[0] = sum[0] + (system.MMinv[i][0]*a.fragments[i])%redundant
		sum[1] = sum[1] + (system.MMinv[i][1]*a.fragments[i])%10
	}
	sum[0] = sum[0] % redundant
	sum[1] = sum[1] % extension

	multiplier := (system.MmodRedInv * (sum[0] - a.fragments[len(system.mods)-2])) % system.mods[len(system.mods)-2]
	for multiplier < 0 {
		multiplier = multiplier + system.mods[len(system.mods)-2]
	}
	a.fragments[len(system.mods)-1] = sum[1] - multiplier*system.MmodRedAnd10[1]
	for a.fragments[len(system.mods)-1] < 0 {
		a.fragments[len(system.mods)-1] = a.fragments[len(system.mods)-1] + 10
	}
	return a
}

// BinaryToRNS converts binary numbers in RNS
func (system System) BinaryToRNS(num float64) *RNS {
	var fragments []int64
	precision := 0

	var sign int
	if num >= 0 {
		sign = 1
	} else {
		sign = -1
		num = num * -1
	}

	_, fractional := math.Modf(num)
	strnum := fmt.Sprintf("%.5f", fractional)
	i := len(strnum) - 1
	for string(strnum[i]) == strconv.Itoa(0) {
		i = i - 1
	}

	precision = -1 * (i - 1)
	correctedNum := num * math.Pow10(i-1)

	for _, mod := range system.mods {
		fragments = append(fragments, int64(correctedNum)%int64(mod))
	}

	return &RNS{
		sign,
		fragments,
		precision,
	}
}

// Multiply returns the multiplication of two RNS numbers
func (system System) Multiply(a, b *RNS) *RNS {
	var fragments []int64
	for index := range a.fragments {
		fragments = append(fragments, (a.fragments[index]*b.fragments[index])%system.mods[index])
	}

	sign := a.sign * b.sign

	return &RNS{
		sign,
		fragments,
		a.precision + b.precision,
	}
}

// Add returns the addition of two RNS numbers
func (system System) Add(a, b *RNS) *RNS {
	var fragments []int64
	if a.precision == b.precision {
		for index := range a.fragments {
			fragments = append(fragments, (a.fragments[index]+b.fragments[index])%system.mods[index])
		}
		return &RNS{
			1,
			fragments,
			a.precision,
		}
	}

	return nil
}

// MultiplicativeInverse finds the multiplicative inverse of a number
func (system System) MultiplicativeInverse(a *RNS) *RNS {
	var fragments []int64
	for index, mod := range system.mods {
		var multiplier = int64(0)
		for (a.fragments[index]*multiplier)%mod != 1 {
			multiplier = multiplier + 1
		}
		fragments = append(fragments, multiplier)
	}
	return &RNS{
		a.sign,
		fragments,
		a.precision,
	}
}

// AdditiveInverse finds the additive inverse of a number
func (system System) AdditiveInverse(a *RNS) *RNS {
	var fragments []int64
	for index, mod := range system.mods {
		var adder = int64(0)
		for (a.fragments[index]+adder)%mod != 0 {
			adder = adder + 1
		}
		fragments = append(fragments, adder)
	}
	return &RNS{
		a.sign,
		fragments,
		a.precision,
	}
}

// Subtract returns the difference between two RNS numbers
func (system System) Subtract(a, b *RNS) *RNS {
	bInv := system.AdditiveInverse(b)
	difference := system.Add(a, bInv)
	difference.precision = a.precision
	return difference
}

// Divide returns the division of two RNS numbers
func (system System) Divide(a, b *RNS) *RNS {
	bInv := system.MultiplicativeInverse(b)
	quotient := system.Multiply(a, bInv)
	quotient.precision = a.precision - b.precision
	return quotient
}

// RNSToBinary returns the binary form of a RNS
func (system System) RNSToBinary(a *RNS) float64 {
	var number = int64(0)
	for index, mod := range system.mods[:len(system.mods)-2] {
		var multiplier = int64(1)
		Mi := (system.M / mod) % mod
		for (Mi*multiplier)%mod != 1 {
			multiplier = multiplier + 1
		}
		number = number + a.fragments[index]*(system.M/mod)*multiplier
	}
	number = number % system.M

	return float64(a.sign) * float64(number) * math.Pow10(a.precision)
}
