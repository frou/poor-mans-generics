package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"github.com/frou/stdext"
)

type Int32Set map[int32]stdext.Unit

func NewInt32Set(initialElements ...int32) Int32Set {
	set := make(Int32Set)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Int32Set) Add(x int32) {
	set[x] = stdext.Extant
}

func (set Int32Set) Remove(x int32) {
	delete(set, x)
}

func (set Int32Set) Contains(x int32) bool {
	_, ok := set[x]
	return ok
}

func (set Int32Set) Is(vals ...int32) bool {
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

func (set Int32Set) Count() int {
	return len(set)
}

func (set Int32Set) Elements() []int32 {
	elm := make([]int32, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Int32Set) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Int32Set) Clone() Int32Set {
	result := NewInt32Set()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Int32Set) Union(b Int32Set) Int32Set {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Int32Set) Intersection(b Int32Set) Int32Set {
	result := NewInt32Set()
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

func (a Int32Set) Difference(b Int32Set) Int32Set {
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
func (a Int32Set) Subtract(b Int32Set) Int32Set {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
