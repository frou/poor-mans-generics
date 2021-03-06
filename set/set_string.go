// Code generated by github.com/frou/poor-mans-generics ; DO NOT EDIT.

package set

type Strings struct {
	backing map[string]struct{}
}

func NewStrings(initialElements ...string) *Strings {
	set := new(Strings)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set *Strings) Add(x string) {
	if set.backing == nil {
		set.backing = make(map[string]struct{})
	}
	set.backing[x] = struct{}{}
}

func (set *Strings) Remove(x string) {
	delete(set.backing, x)
}

func (set *Strings) Contains(x string) bool {
	_, ok := set.backing[x]
	return ok
}

func (set *Strings) Comprises(vals ...string) bool {
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

func (set *Strings) Count() int {
	return len(set.backing)
}

func (set *Strings) Elements() []string {
	elm := make([]string, 0, set.Count())
	for x, _ := range set.backing {
		elm = append(elm, x)
	}
	return elm
}

func (set *Strings) Clear() {
	for x, _ := range set.backing {
		set.Remove(x)
	}
}

func (set *Strings) Clone() *Strings {
	result := NewStrings()
	for x, _ := range set.backing {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a *Strings) Union(b *Strings) *Strings {
	result := a.Clone()
	for x, _ := range b.backing {
		result.Add(x)
	}
	return result
}

func (a *Strings) Intersection(b *Strings) *Strings {
	result := NewStrings()

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

func (a *Strings) Difference(b *Strings) *Strings {
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
func (a *Strings) Subtract(b *Strings) *Strings {
	result := a.Clone()
	for x, _ := range b.backing {
		result.Remove(x)
	}
	return result
}
