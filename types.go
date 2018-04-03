package rns

// System will contain all the member functions of RNS
type System struct {
	mods         []int64
	M            int64
	MMinv        [][]int64
	MmodRedInv   int64
	MmodRedAnd10 []int64
}

// RNS defines the data type of the residue number system
type RNS struct {
	sign      int
	fragments []int64
	precision int
}