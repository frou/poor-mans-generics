package maths

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func MinUint64(x, y uint64) uint64 {
	if x < y {
		return x
	}
	return y
}

func MaxUint64(x, y uint64) uint64 {
	if x > y {
		return x
	}
	return y
}

func AbsUint64(x uint64) uint64 {
	if x < 0 {
		return -x
	}
	return x
}
