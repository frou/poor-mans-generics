package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"github.com/frou/stdext"
)

type Ints map[int]stdext.Unit

func NewInts(initialElements ...int) Ints {
	set := make(Ints)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Ints) Add(x int) {
	set[x] = stdext.Extant
}

func (set Ints) Remove(x int) {
	delete(set, x)
}

func (set Ints) Contains(x int) bool {
	_, ok := set[x]
	return ok
}

func (set Ints) Comprises(vals ...int) bool {
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

func (set Ints) Count() int {
	return len(set)
}

func (set Ints) Elements() []int {
	elm := make([]int, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Ints) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Ints) Clone() Ints {
	result := NewInts()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Ints) Union(b Ints) Ints {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Ints) Intersection(b Ints) Ints {
	result := NewInts()

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

func (a Ints) Difference(b Ints) Ints {
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
func (a Ints) Subtract(b Ints) Ints {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
