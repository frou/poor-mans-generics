// Code generated by github.com/frou/poor-mans-generics ; DO NOT EDIT.

package maths

func MinInt32(x, y int32) int32 {
	if x < y {
		return x
	}
	return y
}

func MaxInt32(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}

func AbsInt32(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}
