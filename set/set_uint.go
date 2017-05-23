package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"github.com/frou/stdext"
)

type Uints map[uint]stdext.Unit

func NewUints(initialElements ...uint) Uints {
	set := make(Uints)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Uints) Add(x uint) {
	set[x] = stdext.Extant
}

func (set Uints) Remove(x uint) {
	delete(set, x)
}

func (set Uints) Contains(x uint) bool {
	_, ok := set[x]
	return ok
}

func (set Uints) Comprises(vals ...uint) bool {
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

func (set Uints) Count() int {
	return len(set)
}

func (set Uints) Elements() []uint {
	elm := make([]uint, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Uints) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Uints) Clone() Uints {
	result := NewUints()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Uints) Union(b Uints) Uints {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Uints) Intersection(b Uints) Uints {
	result := NewUints()
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

func (a Uints) Difference(b Uints) Uints {
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
func (a Uints) Subtract(b Uints) Uints {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
