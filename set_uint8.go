package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"github.com/frou/stdext"
)

type Uint8Set map[uint8]stdext.Unit

func NewUint8Set(initialElements ...uint8) Uint8Set {
	set := make(Uint8Set)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Uint8Set) Add(x uint8) {
	set[x] = stdext.Extant
}

func (set Uint8Set) Remove(x uint8) {
	delete(set, x)
}

func (set Uint8Set) Contains(x uint8) bool {
	_, ok := set[x]
	return ok
}

func (set Uint8Set) Is(vals ...uint8) bool {
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

func (set Uint8Set) Count() int {
	return len(set)
}

func (set Uint8Set) Elements() []uint8 {
	elm := make([]uint8, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Uint8Set) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Uint8Set) Clone() Uint8Set {
	result := NewUint8Set()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Uint8Set) Union(b Uint8Set) Uint8Set {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Uint8Set) Intersection(b Uint8Set) Uint8Set {
	result := NewUint8Set()
	for as, _ := range a {
		for bs, _ := range b {
			if as == bs {
				result.Add(as)
				break
			}
		}
	}
	return result
}

func (a Uint8Set) Difference(b Uint8Set) Uint8Set {
	result := a.Union(b)
	for as, _ := range a {
		for bs, _ := range b {
			if as == bs {
				result.Remove(as)
				break
			}
		}
	}
	return result
}

// Subtraction is non-commutative: a-b is different to b-a.
func (a Uint8Set) Subtract(b Uint8Set) Uint8Set {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
