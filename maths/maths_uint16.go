package maths

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func MinUint16(x, y uint16) uint16 {
	if x < y {
		return x
	}
	return y
}

func MaxUint16(x, y uint16) uint16 {
	if x > y {
		return x
	}
	return y
}

func AbsUint16(x uint16) uint16 {
	if x < 0 {
		return -x
	}
	return x
}
