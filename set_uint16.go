package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"github.com/frou/stdext"
)

type Uint16Set map[uint16]stdext.Unit

func NewUint16Set(initialElements ...uint16) Uint16Set {
	set := make(Uint16Set)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Uint16Set) Add(x uint16) {
	set[x] = stdext.Extant
}

func (set Uint16Set) Remove(x uint16) {
	delete(set, x)
}

func (set Uint16Set) Contains(x uint16) bool {
	_, ok := set[x]
	return ok
}

func (set Uint16Set) Is(vals ...uint16) bool {
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

func (set Uint16Set) Count() int {
	return len(set)
}

func (set Uint16Set) Elements() []uint16 {
	elm := make([]uint16, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Uint16Set) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Uint16Set) Clone() Uint16Set {
	result := NewUint16Set()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Uint16Set) Union(b Uint16Set) Uint16Set {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Uint16Set) Intersection(b Uint16Set) Uint16Set {
	result := NewUint16Set()
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

func (a Uint16Set) Difference(b Uint16Set) Uint16Set {
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
func (a Uint16Set) Subtract(b Uint16Set) Uint16Set {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
