package maths

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func MinUint(x, y uint) uint {
	if x < y {
		return x
	}
	return y
}

func MaxUint(x, y uint) uint {
	if x > y {
		return x
	}
	return y
}

func AbsUint(x uint) uint {
	if x < 0 {
		return -x
	}
	return x
}
