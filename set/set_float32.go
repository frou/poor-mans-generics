package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"github.com/frou/stdext"
)

type Float32s map[float32]stdext.Unit

func NewFloat32s(initialElements ...float32) Float32s {
	set := make(Float32s)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Float32s) Add(x float32) {
	set[x] = stdext.Extant
}

func (set Float32s) Remove(x float32) {
	delete(set, x)
}

func (set Float32s) Contains(x float32) bool {
	_, ok := set[x]
	return ok
}

func (set Float32s) Comprises(vals ...float32) bool {
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

func (set Float32s) Count() int {
	return len(set)
}

func (set Float32s) Elements() []float32 {
	elm := make([]float32, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Float32s) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Float32s) Clone() Float32s {
	result := NewFloat32s()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Float32s) Union(b Float32s) Float32s {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Float32s) Intersection(b Float32s) Float32s {
	result := NewFloat32s()
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

func (a Float32s) Difference(b Float32s) Float32s {
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
func (a Float32s) Subtract(b Float32s) Float32s {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}