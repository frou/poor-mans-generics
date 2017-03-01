package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"github.com/frou/stdext"
)

type Float64Set map[float64]stdext.Unit

func NewFloat64Set(initialElements ...float64) Float64Set {
	set := make(Float64Set)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Float64Set) Add(x float64) {
	set[x] = stdext.Extant
}

func (set Float64Set) Remove(x float64) {
	delete(set, x)
}

func (set Float64Set) Contains(x float64) bool {
	_, ok := set[x]
	return ok
}

func (set Float64Set) Comprises(vals ...float64) bool {
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

func (set Float64Set) Count() int {
	return len(set)
}

func (set Float64Set) Elements() []float64 {
	elm := make([]float64, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Float64Set) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Float64Set) Clone() Float64Set {
	result := NewFloat64Set()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Float64Set) Union(b Float64Set) Float64Set {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Float64Set) Intersection(b Float64Set) Float64Set {
	result := NewFloat64Set()
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

func (a Float64Set) Difference(b Float64Set) Float64Set {
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
func (a Float64Set) Subtract(b Float64Set) Float64Set {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
