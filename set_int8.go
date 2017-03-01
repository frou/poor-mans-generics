package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"github.com/frou/stdext"
)

type Int8Set map[int8]stdext.Unit

func NewInt8Set(initialElements ...int8) Int8Set {
	set := make(Int8Set)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Int8Set) Add(x int8) {
	set[x] = stdext.Extant
}

func (set Int8Set) Remove(x int8) {
	delete(set, x)
}

func (set Int8Set) Contains(x int8) bool {
	_, ok := set[x]
	return ok
}

func (set Int8Set) Comprises(vals ...int8) bool {
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

func (set Int8Set) Count() int {
	return len(set)
}

func (set Int8Set) Elements() []int8 {
	elm := make([]int8, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Int8Set) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Int8Set) Clone() Int8Set {
	result := NewInt8Set()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Int8Set) Union(b Int8Set) Int8Set {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Int8Set) Intersection(b Int8Set) Int8Set {
	result := NewInt8Set()
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

func (a Int8Set) Difference(b Int8Set) Int8Set {
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
func (a Int8Set) Subtract(b Int8Set) Int8Set {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
