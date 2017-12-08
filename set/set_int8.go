package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

type Int8s map[int8]struct{}

func NewInt8s(initialElements ...int8) Int8s {
	set := make(Int8s)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Int8s) Add(x int8) {
	set[x] = struct{}{}
}

func (set Int8s) Remove(x int8) {
	delete(set, x)
}

func (set Int8s) Contains(x int8) bool {
	_, ok := set[x]
	return ok
}

func (set Int8s) Comprises(vals ...int8) bool {
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

func (set Int8s) Count() int {
	return len(set)
}

func (set Int8s) Elements() []int8 {
	elm := make([]int8, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Int8s) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Int8s) Clone() Int8s {
	result := NewInt8s()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Int8s) Union(b Int8s) Int8s {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Int8s) Intersection(b Int8s) Int8s {
	result := NewInt8s()

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

func (a Int8s) Difference(b Int8s) Int8s {
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
func (a Int8s) Subtract(b Int8s) Int8s {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
