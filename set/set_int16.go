package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"github.com/frou/stdext"
)

type Int16s map[int16]stdext.Unit

func NewInt16s(initialElements ...int16) Int16s {
	set := make(Int16s)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Int16s) Add(x int16) {
	set[x] = stdext.Extant
}

func (set Int16s) Remove(x int16) {
	delete(set, x)
}

func (set Int16s) Contains(x int16) bool {
	_, ok := set[x]
	return ok
}

func (set Int16s) Comprises(vals ...int16) bool {
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

func (set Int16s) Count() int {
	return len(set)
}

func (set Int16s) Elements() []int16 {
	elm := make([]int16, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Int16s) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Int16s) Clone() Int16s {
	result := NewInt16s()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Int16s) Union(b Int16s) Int16s {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Int16s) Intersection(b Int16s) Int16s {
	result := NewInt16s()

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

func (a Int16s) Difference(b Int16s) Int16s {
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
func (a Int16s) Subtract(b Int16s) Int16s {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
