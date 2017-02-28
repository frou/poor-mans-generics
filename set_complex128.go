package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"github.com/frou/stdext"
)

type Complex128Set map[complex128]stdext.Unit

func NewComplex128Set(initialElements ...complex128) Complex128Set {
	set := make(Complex128Set)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Complex128Set) Add(x complex128) {
	set[x] = stdext.Extant
}

func (set Complex128Set) Remove(x complex128) {
	delete(set, x)
}

func (set Complex128Set) Contains(x complex128) bool {
	_, ok := set[x]
	return ok
}

func (set Complex128Set) Is(vals ...complex128) bool {
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

func (set Complex128Set) Count() int {
	return len(set)
}

func (set Complex128Set) Elements() []complex128 {
	elm := make([]complex128, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Complex128Set) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Complex128Set) Clone() Complex128Set {
	result := NewComplex128Set()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Complex128Set) Union(b Complex128Set) Complex128Set {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Complex128Set) Intersection(b Complex128Set) Complex128Set {
	result := NewComplex128Set()
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

func (a Complex128Set) Difference(b Complex128Set) Complex128Set {
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
func (a Complex128Set) Subtract(b Complex128Set) Complex128Set {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
