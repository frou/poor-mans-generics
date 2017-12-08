package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

type Uint32s map[uint32]struct{}

func NewUint32s(initialElements ...uint32) Uint32s {
	set := make(Uint32s)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Uint32s) Add(x uint32) {
	set[x] = struct{}{}
}

func (set Uint32s) Remove(x uint32) {
	delete(set, x)
}

func (set Uint32s) Contains(x uint32) bool {
	_, ok := set[x]
	return ok
}

func (set Uint32s) Comprises(vals ...uint32) bool {
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

func (set Uint32s) Count() int {
	return len(set)
}

func (set Uint32s) Elements() []uint32 {
	elm := make([]uint32, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Uint32s) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Uint32s) Clone() Uint32s {
	result := NewUint32s()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Uint32s) Union(b Uint32s) Uint32s {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Uint32s) Intersection(b Uint32s) Uint32s {
	result := NewUint32s()

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

func (a Uint32s) Difference(b Uint32s) Uint32s {
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
func (a Uint32s) Subtract(b Uint32s) Uint32s {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
