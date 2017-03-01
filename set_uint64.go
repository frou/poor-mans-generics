package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"github.com/frou/stdext"
)

type Uint64s map[uint64]stdext.Unit

func NewUint64s(initialElements ...uint64) Uint64s {
	set := make(Uint64s)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Uint64s) Add(x uint64) {
	set[x] = stdext.Extant
}

func (set Uint64s) Remove(x uint64) {
	delete(set, x)
}

func (set Uint64s) Contains(x uint64) bool {
	_, ok := set[x]
	return ok
}

func (set Uint64s) Comprises(vals ...uint64) bool {
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

func (set Uint64s) Count() int {
	return len(set)
}

func (set Uint64s) Elements() []uint64 {
	elm := make([]uint64, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Uint64s) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Uint64s) Clone() Uint64s {
	result := NewUint64s()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Uint64s) Union(b Uint64s) Uint64s {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Uint64s) Intersection(b Uint64s) Uint64s {
	result := NewUint64s()
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

func (a Uint64s) Difference(b Uint64s) Uint64s {
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
func (a Uint64s) Subtract(b Uint64s) Uint64s {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
