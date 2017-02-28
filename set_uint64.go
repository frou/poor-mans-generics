package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"github.com/frou/stdext"
)

type Uint64Set map[uint64]stdext.Unit

func NewUint64Set(initialElements ...uint64) Uint64Set {
	set := make(Uint64Set)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Uint64Set) Add(x uint64) {
	set[x] = stdext.Extant
}

func (set Uint64Set) Remove(x uint64) {
	delete(set, x)
}

func (set Uint64Set) Contains(x uint64) bool {
	_, ok := set[x]
	return ok
}

func (set Uint64Set) Is(vals ...uint64) bool {
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

func (set Uint64Set) Count() int {
	return len(set)
}

func (set Uint64Set) Elements() []uint64 {
	elm := make([]uint64, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Uint64Set) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Uint64Set) Clone() Uint64Set {
	result := NewUint64Set()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Uint64Set) Union(b Uint64Set) Uint64Set {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Uint64Set) Intersection(b Uint64Set) Uint64Set {
	result := NewUint64Set()
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

func (a Uint64Set) Difference(b Uint64Set) Uint64Set {
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
func (a Uint64Set) Subtract(b Uint64Set) Uint64Set {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
