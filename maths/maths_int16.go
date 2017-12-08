package maths

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func MinInt16(x, y int16) int16 {
	if x < y {
		return x
	}
	return y
}

func MaxInt16(x, y int16) int16 {
	if x > y {
		return x
	}
	return y
}

func AbsInt16(x int16) int16 {
	if x < 0 {
		return -x
	}
	return x
}
