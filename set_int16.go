package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"github.com/frou/stdext"
)

type Int16Set map[int16]stdext.Unit

func NewInt16Set(initialElements ...int16) Int16Set {
	set := make(Int16Set)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Int16Set) Add(x int16) {
	set[x] = stdext.Extant
}

func (set Int16Set) Remove(x int16) {
	delete(set, x)
}

func (set Int16Set) Contains(x int16) bool {
	_, ok := set[x]
	return ok
}

func (set Int16Set) Comprises(vals ...int16) bool {
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

func (set Int16Set) Count() int {
	return len(set)
}

func (set Int16Set) Elements() []int16 {
	elm := make([]int16, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Int16Set) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Int16Set) Clone() Int16Set {
	result := NewInt16Set()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Int16Set) Union(b Int16Set) Int16Set {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Int16Set) Intersection(b Int16Set) Int16Set {
	result := NewInt16Set()
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

func (a Int16Set) Difference(b Int16Set) Int16Set {
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
func (a Int16Set) Subtract(b Int16Set) Int16Set {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
