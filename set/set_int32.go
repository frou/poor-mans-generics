package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

type Int32s map[int32]struct{}

func NewInt32s(initialElements ...int32) Int32s {
	set := make(Int32s)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Int32s) Add(x int32) {
	set[x] = struct{}{}
}

func (set Int32s) Remove(x int32) {
	delete(set, x)
}

func (set Int32s) Contains(x int32) bool {
	_, ok := set[x]
	return ok
}

func (set Int32s) Comprises(vals ...int32) bool {
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

func (set Int32s) Count() int {
	return len(set)
}

func (set Int32s) Elements() []int32 {
	elm := make([]int32, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Int32s) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Int32s) Clone() Int32s {
	result := NewInt32s()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Int32s) Union(b Int32s) Int32s {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Int32s) Intersection(b Int32s) Int32s {
	result := NewInt32s()

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

func (a Int32s) Difference(b Int32s) Int32s {
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
func (a Int32s) Subtract(b Int32s) Int32s {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
