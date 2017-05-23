package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"github.com/frou/stdext"
)

type Uint16s map[uint16]stdext.Unit

func NewUint16s(initialElements ...uint16) Uint16s {
	set := make(Uint16s)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Uint16s) Add(x uint16) {
	set[x] = stdext.Extant
}

func (set Uint16s) Remove(x uint16) {
	delete(set, x)
}

func (set Uint16s) Contains(x uint16) bool {
	_, ok := set[x]
	return ok
}

func (set Uint16s) Comprises(vals ...uint16) bool {
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

func (set Uint16s) Count() int {
	return len(set)
}

func (set Uint16s) Elements() []uint16 {
	elm := make([]uint16, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Uint16s) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Uint16s) Clone() Uint16s {
	result := NewUint16s()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Uint16s) Union(b Uint16s) Uint16s {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Uint16s) Intersection(b Uint16s) Uint16s {
	result := NewUint16s()
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

func (a Uint16s) Difference(b Uint16s) Uint16s {
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
func (a Uint16s) Subtract(b Uint16s) Uint16s {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}