package maths

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func MinUint32(x, y uint32) uint32 {
	if x < y {
		return x
	}
	return y
}

func MaxUint32(x, y uint32) uint32 {
	if x > y {
		return x
	}
	return y
}

func AbsUint32(x uint32) uint32 {
	if x < 0 {
		return -x
	}
	return x
}
