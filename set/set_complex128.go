package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

type Complex128s map[complex128]struct{}

func NewComplex128s(initialElements ...complex128) Complex128s {
	set := make(Complex128s)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Complex128s) Add(x complex128) {
	set[x] = struct{}{}
}

func (set Complex128s) Remove(x complex128) {
	delete(set, x)
}

func (set Complex128s) Contains(x complex128) bool {
	_, ok := set[x]
	return ok
}

func (set Complex128s) Comprises(vals ...complex128) bool {
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

func (set Complex128s) Count() int {
	return len(set)
}

func (set Complex128s) Elements() []complex128 {
	elm := make([]complex128, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Complex128s) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Complex128s) Clone() Complex128s {
	result := NewComplex128s()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Complex128s) Union(b Complex128s) Complex128s {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Complex128s) Intersection(b Complex128s) Complex128s {
	result := NewComplex128s()

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

func (a Complex128s) Difference(b Complex128s) Complex128s {
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
func (a Complex128s) Subtract(b Complex128s) Complex128s {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
