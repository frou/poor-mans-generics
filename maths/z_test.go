package maths

import (
	"testing"
)

// `go generate` needs to have been run before tests can be run.

func TestMin(t *testing.T) {
	var a, b int16 = 3, 5
	if MinInt16(a, b) != a {
		t.Fail()
	}

	var c, d int16 = 5, -7
	if MinInt16(c, d) != d {
		t.Fail()
	}
}

func TestMax(t *testing.T) {
	var a, b int64 = 333, 55555
	if MaxInt64(a, b) != b {
		t.Fail()
	}

	var c, d int64 = 1, -7777777
	if MaxInt64(c, d) != c {
		t.Fail()
	}
}

// Abs* functions are generated for unsigned integer types, too, which is
// pointless. But it doesn't hurt anything.

func TestAbs(t *testing.T) {
	var a int8 = 8
	if AbsInt8(a) != a {
		t.Fail()
	}

	var b int8 = -20
	if AbsInt8(b) != 20 {
		t.Fail()
	}
}
