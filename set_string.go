package set

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// !! This code was automatically generated. Do not hand edit.
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"github.com/frou/stdext"
)

type StringSet map[string]stdext.Unit

func NewStringSet(initialElements ...string) StringSet {
	set := make(StringSet)
	for _, x := range initialElements {
		set.Add(x)
	}
	return set
}

func (set StringSet) Add(x string) {
	set[x] = stdext.Extant
}

func (set StringSet) Remove(x string) {
	delete(set, x)
}

func (set StringSet) Contains(x string) bool {
	_, ok := set[x]
	return ok
}

func (set StringSet) Is(vals ...string) bool {
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

func (set StringSet) Count() int {
	return len(set)
}

func (set StringSet) Elements() []string {
	elm := make([]string, 0, set.Count())
	for x, _ := range set {
		elm = append(elm, x)
	}
	return elm
}

func (set StringSet) Clear() {
	for x, _ := range set {
		set.Remove(x)
	}
}

func (set StringSet) Clone() StringSet {
	result := NewStringSet()
	for x, _ := range set {
		result.Add(x)
	}
	return result
}

// ------------------------------------------------------------

func (a StringSet) Union(b StringSet) StringSet {
	result := a.Clone()
	for x, _ := range b {
		result.Add(x)
	}
	return result
}

func (a StringSet) Intersection(b StringSet) StringSet {
	result := NewStringSet()
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

func (a StringSet) Difference(b StringSet) StringSet {
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
func (a StringSet) Subtract(b StringSet) StringSet {
	result := a.Clone()
	for x, _ := range b {
		result.Remove(x)
	}
	return result
}
