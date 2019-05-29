package rnums

// NewSystemFromMods returns the system object when mods are given to it
func NewSystemFromMods(mods []int64, redundant int64, extension int64) System {
	var M = int64(1)
	for _, mod := range mods {
		M = M * mod
	}

	var MMinv [][]int64
	for _, mod := range mods {
		Mi := M / mod
		modMiRed := Mi % redundant
		modMiExt := Mi % extension
		MiMod := Mi % mod
		var multiplier = int64(0)
		for (MiMod*multiplier)%mod != 1 {
			multiplier = multiplier + 1
		}
		modMiRedInv := multiplier % redundant
		modMiExtInv := multiplier % extension
		MMinv = append(MMinv, []int64{(modMiRed * modMiRedInv) % redundant, (modMiExt * modMiExtInv) % extension})
	}

	mods = append(mods, redundant)
	mods = append(mods, extension)

	MmodRedAndExt := []int64{M % redundant, M % extension}

	MmodRed := M % redundant
	MmodRedInv := int64(0)
	for (MmodRed*MmodRedInv)%redundant != 1 {
		MmodRedInv = MmodRedInv + 1
	}

	return System{
		mods,
		M,
		MMinv,
		MmodRedInv,
		MmodRedAndExt,
	}
}

// NewSystem returns a pointer to an object of type System
func NewSystem(noOfFragments int, redundant int64, extension int64) System {
	var mods []int64
	var num, flag = 7, 0

	for noOfFragments > 0 {
		flag = 0
		for i := 2; i < num/2; i++ {
			if num%i == 0 {
				flag = 1
				break
			}
		}
		if flag != 1 {
			mods = append(mods, int64(num))
			noOfFragments = noOfFragments - 1
		}
		num = num + 2
	}

	return NewSystemFromMods(mods, redundant, extension)
}
