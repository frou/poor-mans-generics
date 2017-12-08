package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

type Float64s map[float64]struct{}

func NewFloat64s(initialElements ...float64) Float64s {
	set := make(Float64s)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Float64s) Add(x float64) {
	set[x] = struct{}{}
}

func (set Float64s) Remove(x float64) {
	delete(set, x)
}

func (set Float64s) Contains(x float64) bool {
	_, ok := set[x]
	return ok
}

func (set Float64s) Comprises(vals ...float64) bool {
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

func (set Float64s) Count() int {
	return len(set)
}

func (set Float64s) Elements() []float64 {
	elm := make([]float64, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Float64s) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Float64s) Clone() Float64s {
	result := NewFloat64s()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Float64s) Union(b Float64s) Float64s {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Float64s) Intersection(b Float64s) Float64s {
	result := NewFloat64s()

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

func (a Float64s) Difference(b Float64s) Float64s {
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
func (a Float64s) Subtract(b Float64s) Float64s {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
