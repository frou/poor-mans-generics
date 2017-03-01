package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"github.com/frou/stdext"
)

type Float32Set map[float32]stdext.Unit

func NewFloat32Set(initialElements ...float32) Float32Set {
	set := make(Float32Set)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Float32Set) Add(x float32) {
	set[x] = stdext.Extant
}

func (set Float32Set) Remove(x float32) {
	delete(set, x)
}

func (set Float32Set) Contains(x float32) bool {
	_, ok := set[x]
	return ok
}

func (set Float32Set) Comprises(vals ...float32) bool {
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

func (set Float32Set) Count() int {
	return len(set)
}

func (set Float32Set) Elements() []float32 {
	elm := make([]float32, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Float32Set) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Float32Set) Clone() Float32Set {
	result := NewFloat32Set()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Float32Set) Union(b Float32Set) Float32Set {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Float32Set) Intersection(b Float32Set) Float32Set {
	result := NewFloat32Set()
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

func (a Float32Set) Difference(b Float32Set) Float32Set {
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
func (a Float32Set) Subtract(b Float32Set) Float32Set {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
