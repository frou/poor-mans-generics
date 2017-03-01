package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"github.com/frou/stdext"
)

type Int64Set map[int64]stdext.Unit

func NewInt64Set(initialElements ...int64) Int64Set {
	set := make(Int64Set)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Int64Set) Add(x int64) {
	set[x] = stdext.Extant
}

func (set Int64Set) Remove(x int64) {
	delete(set, x)
}

func (set Int64Set) Contains(x int64) bool {
	_, ok := set[x]
	return ok
}

func (set Int64Set) Comprises(vals ...int64) bool {
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

func (set Int64Set) Count() int {
	return len(set)
}

func (set Int64Set) Elements() []int64 {
	elm := make([]int64, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Int64Set) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Int64Set) Clone() Int64Set {
	result := NewInt64Set()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Int64Set) Union(b Int64Set) Int64Set {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Int64Set) Intersection(b Int64Set) Int64Set {
	result := NewInt64Set()
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

func (a Int64Set) Difference(b Int64Set) Int64Set {
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
func (a Int64Set) Subtract(b Int64Set) Int64Set {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
