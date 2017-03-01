package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"github.com/frou/stdext"
)

type Complex64Set map[complex64]stdext.Unit

func NewComplex64Set(initialElements ...complex64) Complex64Set {
	set := make(Complex64Set)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Complex64Set) Add(x complex64) {
	set[x] = stdext.Extant
}

func (set Complex64Set) Remove(x complex64) {
	delete(set, x)
}

func (set Complex64Set) Contains(x complex64) bool {
	_, ok := set[x]
	return ok
}

func (set Complex64Set) Comprises(vals ...complex64) bool {
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

func (set Complex64Set) Count() int {
	return len(set)
}

func (set Complex64Set) Elements() []complex64 {
	elm := make([]complex64, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Complex64Set) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Complex64Set) Clone() Complex64Set {
	result := NewComplex64Set()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Complex64Set) Union(b Complex64Set) Complex64Set {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Complex64Set) Intersection(b Complex64Set) Complex64Set {
	result := NewComplex64Set()
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

func (a Complex64Set) Difference(b Complex64Set) Complex64Set {
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
func (a Complex64Set) Subtract(b Complex64Set) Complex64Set {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
