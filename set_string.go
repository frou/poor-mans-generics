package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"github.com/frou/stdext"
)

type Strings map[string]stdext.Unit

func NewStrings(initialElements ...string) Strings {
	set := make(Strings)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set Strings) Add(x string) {
	set[x] = stdext.Extant
}

func (set Strings) Remove(x string) {
	delete(set, x)
}

func (set Strings) Contains(x string) bool {
	_, ok := set[x]
	return ok
}

func (set Strings) Comprises(vals ...string) bool {
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

func (set Strings) Count() int {
	return len(set)
}

func (set Strings) Elements() []string {
	elm := make([]string, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set Strings) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set Strings) Clone() Strings {
	result := NewStrings()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a Strings) Union(b Strings) Strings {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a Strings) Intersection(b Strings) Strings {
	result := NewStrings()
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

func (a Strings) Difference(b Strings) Strings {
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
func (a Strings) Subtract(b Strings) Strings {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
