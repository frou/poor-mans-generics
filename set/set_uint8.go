package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"github.com/frou/stdext"
)

type Uint8s map[uint8]stdext.Unit

func NewUint8s(initialElements ...uint8) Uint8s {
	set := make(Uint8s)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Uint8s) Add(x uint8) {
	set[x] = stdext.Extant
}

func (set Uint8s) Remove(x uint8) {
	delete(set, x)
}

func (set Uint8s) Contains(x uint8) bool {
	_, ok := set[x]
	return ok
}

func (set Uint8s) Comprises(vals ...uint8) bool {
	if set.Count() != len(vals) {
		return false
	}
	for _, v := range vals {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

func (set Uint8s) Count() int {
	return len(set)
}

func (set Uint8s) Elements() []uint8 {
	elm := make([]uint8, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Uint8s) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Uint8s) Clone() Uint8s {
	result := NewUint8s()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Uint8s) Union(b Uint8s) Uint8s {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Uint8s) Intersection(b Uint8s) Uint8s {
	result := NewUint8s()

	smaller, larger := a, b
	if a.Count() > b.Count() {
		smaller, larger = b, a
	}

	for x, _ := range smaller {
		if larger.Contains(x) {
			result.Add(x)
		}
	}
	return result
}

func (a Uint8s) Difference(b Uint8s) Uint8s {
	result := a.Union(b)

	smaller, larger := a, b
	if a.Count() > b.Count() {
		smaller, larger = b, a
	}

	for x, _ := range smaller {
		if larger.Contains(x) {
			result.Remove(x)
		}
	}
	return result
}

// Subtraction is non-commutative: a-b is different to b-a.
func (a Uint8s) Subtract(b Uint8s) Uint8s {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
