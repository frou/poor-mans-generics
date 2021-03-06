// Code generated by github.com/frou/poor-mans-generics ; DO NOT EDIT.

package set

type Uint64s struct {
	backing map[uint64]struct{}
}

func NewUint64s(initialElements ...uint64) *Uint64s {
	set := new(Uint64s)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set *Uint64s) Add(x uint64) {
	if set.backing == nil {
		set.backing = make(map[uint64]struct{})
	}
	set.backing[x] = struct{}{}
}

func (set *Uint64s) Remove(x uint64) {
	delete(set.backing, x)
}

func (set *Uint64s) Contains(x uint64) bool {
	_, ok := set.backing[x]
	return ok
}

func (set *Uint64s) Comprises(vals ...uint64) bool {
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

func (set *Uint64s) Count() int {
	return len(set.backing)
}

func (set *Uint64s) Elements() []uint64 {
	elm := make([]uint64, 0, set.Count())
	for x, _ := range set.backing {
		elm = append(elm, x)
	}
	return elm
}

func (set *Uint64s) Clear() {
	for x, _ := range set.backing {
		set.Remove(x)
	}
}

func (set *Uint64s) Clone() *Uint64s {
	result := NewUint64s()
	for x, _ := range set.backing {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a *Uint64s) Union(b *Uint64s) *Uint64s {
	result := a.Clone()
	for x, _ := range b.backing {
		result.Add(x)
	}
	return result
}

func (a *Uint64s) Intersection(b *Uint64s) *Uint64s {
	result := NewUint64s()

	smaller, larger := a, b
	if a.Count() > b.Count() {
		smaller, larger = b, a
	}

	for x, _ := range smaller.backing {
		if larger.Contains(x) {
			result.Add(x)
		}
	}
	return result
}

func (a *Uint64s) Difference(b *Uint64s) *Uint64s {
	result := a.Union(b)

	smaller, larger := a, b
	if a.Count() > b.Count() {
		smaller, larger = b, a
	}

	for x, _ := range smaller.backing {
		if larger.Contains(x) {
			result.Remove(x)
		}
	}
	return result
}

// Subtraction is non-commutative: a-b is different to b-a.
func (a *Uint64s) Subtract(b *Uint64s) *Uint64s {
	result := a.Clone()
	for x, _ := range b.backing {
		result.Remove(x)
	}
	return result
}
