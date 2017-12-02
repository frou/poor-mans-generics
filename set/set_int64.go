package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"github.com/frou/stdext"
)

type Int64s map[int64]stdext.Unit

func NewInt64s(initialElements ...int64) Int64s {
	set := make(Int64s)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Int64s) Add(x int64) {
	set[x] = stdext.Extant
}

func (set Int64s) Remove(x int64) {
	delete(set, x)
}

func (set Int64s) Contains(x int64) bool {
	_, ok := set[x]
	return ok
}

func (set Int64s) Comprises(vals ...int64) bool {
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

func (set Int64s) Count() int {
	return len(set)
}

func (set Int64s) Elements() []int64 {
	elm := make([]int64, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Int64s) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Int64s) Clone() Int64s {
	result := NewInt64s()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Int64s) Union(b Int64s) Int64s {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Int64s) Intersection(b Int64s) Int64s {
	result := NewInt64s()

	smaller, larger := a, b
	if a.Count() > b.Count() {
		smaller, larger = b, a
	}

	for x, _ := range smaller {
		if larger.Contains(x) {
			result.Add(x)
		}
	}
	return result
}

func (a Int64s) Difference(b Int64s) Int64s {
	result := a.Union(b)

	smaller, larger := a, b
	if a.Count() > b.Count() {
		smaller, larger = b, a
	}

	for x, _ := range smaller {
		if larger.Contains(x) {
			result.Remove(x)
		}
	}
	return result
}

// Subtraction is non-commutative: a-b is different to b-a.
func (a Int64s) Subtract(b Int64s) Int64s {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
