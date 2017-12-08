package maths

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func MinFloat32(x, y float32) float32 {
	if x < y {
		return x
	}
	return y
}

func MaxFloat32(x, y float32) float32 {
	if x > y {
		return x
	}
	return y
}

func AbsFloat32(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}
