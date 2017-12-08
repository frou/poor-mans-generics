package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

type Complex64s map[complex64]struct{}

func NewComplex64s(initialElements ...complex64) Complex64s {
	set := make(Complex64s)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Complex64s) Add(x complex64) {
	set[x] = struct{}{}
}

func (set Complex64s) Remove(x complex64) {
	delete(set, x)
}

func (set Complex64s) Contains(x complex64) bool {
	_, ok := set[x]
	return ok
}

func (set Complex64s) Comprises(vals ...complex64) bool {
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

func (set Complex64s) Count() int {
	return len(set)
}

func (set Complex64s) Elements() []complex64 {
	elm := make([]complex64, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Complex64s) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Complex64s) Clone() Complex64s {
	result := NewComplex64s()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Complex64s) Union(b Complex64s) Complex64s {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Complex64s) Intersection(b Complex64s) Complex64s {
	result := NewComplex64s()

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

func (a Complex64s) Difference(b Complex64s) Complex64s {
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
func (a Complex64s) Subtract(b Complex64s) Complex64s {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
