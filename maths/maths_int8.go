package maths

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func MinInt8(x, y int8) int8 {
	if x < y {
		return x
	}
	return y
}

func MaxInt8(x, y int8) int8 {
	if x > y {
		return x
	}
	return y
}

func AbsInt8(x int8) int8 {
	if x < 0 {
		return -x
	}
	return x
}
