package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"github.com/frou/stdext"
)

type Uint32Set map[uint32]stdext.Unit

func NewUint32Set(initialElements ...uint32) Uint32Set {
	set := make(Uint32Set)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Uint32Set) Add(x uint32) {
	set[x] = stdext.Extant
}

func (set Uint32Set) Remove(x uint32) {
	delete(set, x)
}

func (set Uint32Set) Contains(x uint32) bool {
	_, ok := set[x]
	return ok
}

func (set Uint32Set) Is(vals ...uint32) bool {
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

func (set Uint32Set) Count() int {
	return len(set)
}

func (set Uint32Set) Elements() []uint32 {
	elm := make([]uint32, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Uint32Set) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Uint32Set) Clone() Uint32Set {
	result := NewUint32Set()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Uint32Set) Union(b Uint32Set) Uint32Set {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Uint32Set) Intersection(b Uint32Set) Uint32Set {
	result := NewUint32Set()
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

func (a Uint32Set) Difference(b Uint32Set) Uint32Set {
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
func (a Uint32Set) Subtract(b Uint32Set) Uint32Set {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
