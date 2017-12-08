package maths

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func MinUint8(x, y uint8) uint8 {
	if x < y {
		return x
	}
	return y
}

func MaxUint8(x, y uint8) uint8 {
	if x > y {
		return x
	}
	return y
}

func AbsUint8(x uint8) uint8 {
	if x < 0 {
		return -x
	}
	return x
}
