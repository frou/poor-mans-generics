package set

// This assumes a concrete type (StringSet) has been generated.

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet_New_SetIsNotNil(t *testing.T) {
	set := NewStringSet()

	assert.NotNil(t, set)
}

func TestSet_NewWithElements_SetIsPopulated(t *testing.T) {
	set := NewStringSet("a", "b")

	assert.Equal(t, 2, set.Count())
	assert.True(t, set.Is("a", "b"))
}

func TestSet_AddToEmpty_SetIsOneElement(t *testing.T) {
	set := NewStringSet()

	set.Add("a")

	assert.Equal(t, 1, set.Count())
	assert.True(t, set.Is("a"))
}

func TestSet_AddToNonEmpty_SetIsOneMoreElement(t *testing.T) {
	set := NewStringSet()

	set.Add("a")
	set.Add("b")

	assert.Equal(t, 2, set.Count())
	assert.True(t, set.Is("a", "b"))
}

func TestSet_RemoveExtantElement_ElementIsRemoved(t *testing.T) {
	set := NewStringSet()

	set.Add("a")
	set.Add("b")
	assert.Equal(t, 2, set.Count())

	set.Remove("b")
	assert.Equal(t, 1, set.Count())
}

func TestSet_RemoveNonexistentElement_NoEffect(t *testing.T) {
	set := NewStringSet()

	set.Add("a")
	set.Add("b")
	assert.Equal(t, 2, set.Count())

	set.Remove("c")
	assert.Equal(t, 2, set.Count())
}

func TestSet_ContainsAddedElement_Does(t *testing.T) {
	set := NewStringSet()

	set.Add("a")

	assert.True(t, set.Contains("a"))
}

func TestSet_ContainsNonAddedElement_DoesNot(t *testing.T) {
	set := NewStringSet()

	assert.False(t, set.Contains(""))
	set.Add("a")
	assert.False(t, set.Contains("b"))
}

func TestSet_ComprisesExactMatch_Does(t *testing.T) {
	set := NewStringSet()

	assert.True(t, set.Is())
	set.Add("a")
	assert.True(t, set.Is("a"))
	set.Add("b")
	set.Add("c")
	assert.True(t, set.Is("a", "b", "c"))
}

func TestSet_ComprisesInexactMatch_DoesNot(t *testing.T) {
	set := NewStringSet()

	set.Add("a")
	set.Add("b")
	set.Add("c")

	assert.False(t, set.Is())
	assert.False(t, set.Is("a", "b"))
	assert.False(t, set.Is("a", "b", "x"))
	assert.False(t, set.Is("x", "b", "c"))
}

func TestSet_Count_IsAccurate(t *testing.T) {
	set := NewStringSet()

	assert.Zero(t, set.Count())

	set.Add("a")
	assert.Equal(t, 1, set.Count())

	set.Add("b")
	set.Add("c")
	assert.Equal(t, 3, set.Count())

	set.Remove("c")
	assert.Equal(t, 2, set.Count())

	set.Remove("b")
	set.Remove("a")
	assert.Equal(t, 0, set.Count())

	set.Remove("a")
	assert.Equal(t, 0, set.Count())

	set.Remove("x")
	assert.Equal(t, 0, set.Count())
}

func TestSet_ElementsSlice_IsWellFormed(t *testing.T) {
	set := NewStringSet()

	elements := set.Elements()
	assert.Zero(t, len(elements))

	set.Add("a")
	set.Add("b")
	set.Add("c")

	elements = set.Elements()
	assert.Equal(t, 3, len(elements))
	assert.Contains(t, elements, "a")
	assert.Contains(t, elements, "b")
	assert.Contains(t, elements, "c")
}

func TestSet_ElementsSlice_DoesNotAliasSet(t *testing.T) {
	set := NewStringSet()
	set.Add("a")
	set.Add("b")
	set.Add("c")

	elements := set.Elements()
	elements[0] = "x"
	elements[1] = "x"
	elements[2] = "x"

	assert.True(t, set.Is("a", "b", "c"))
}

func TestSet_Clear_SetIsEmpty(t *testing.T) {
	set := NewStringSet()

	set.Add("a")
	set.Add("b")
	set.Add("c")

	set.Clear()
	assert.Zero(t, set.Count())
}

func TestSet_UnionNoneInCommon_ResultIsAll(t *testing.T) {
	l := NewStringSet()
	l.Add("a")
	l.Add("b")
	l.Add("c")

	r := NewStringSet()
	r.Add("x")
	r.Add("y")
	r.Add("z")

	union := l.Union(r)
	assert.Equal(t, 6, union.Count())
	assert.True(t, union.Is("a", "b", "c", "x", "y", "z"))

	// Shouldn't mutate originals
	assert.True(t, l.Is("a", "b", "c"))
	assert.True(t, r.Is("x", "y", "z"))
}

func TestSet_UnionSomeInCommon_ResultIsAll(t *testing.T) {
	l := NewStringSet()
	l.Add("a")
	l.Add("b")
	l.Add("c")

	r := NewStringSet()
	r.Add("d")
	r.Add("b")
	r.Add("c")

	union := l.Union(r)
	assert.Equal(t, 4, union.Count())
	assert.True(t, union.Is("a", "b", "c", "d"))

	// Shouldn't mutate originals
	assert.True(t, l.Is("a", "b", "c"))
	assert.True(t, r.Is("b", "c", "d"))
}

func TestSet_UnionAllInCommon_ResultIsAll(t *testing.T) {
	l := NewStringSet()
	l.Add("a")
	l.Add("b")
	l.Add("c")

	r := NewStringSet()
	r.Add("c")
	r.Add("a")
	r.Add("b")

	union := l.Union(r)
	assert.Equal(t, 3, union.Count())
	assert.True(t, union.Is("a", "b", "c"))

	// Shouldn't mutate originals
	assert.True(t, l.Is("a", "b", "c"))
	assert.True(t, r.Is("a", "b", "c"))
}

func TestSet_IntersectionNoneInCommon_ResultIsNone(t *testing.T) {
	l := NewStringSet()
	l.Add("a")
	l.Add("b")
	l.Add("c")

	r := NewStringSet()
	r.Add("x")
	r.Add("y")
	r.Add("z")

	intersection := l.Intersection(r)
	assert.Equal(t, 0, intersection.Count())
	assert.True(t, intersection.Is())

	// Shouldn't mutate originals
	assert.True(t, l.Is("a", "b", "c"))
	assert.True(t, r.Is("x", "y", "z"))
}

func TestSet_IntersectionSomeInCommon_ResultIsCommon(t *testing.T) {
	l := NewStringSet()
	l.Add("a")
	l.Add("b")
	l.Add("c")

	r := NewStringSet()
	r.Add("b")
	r.Add("c")
	r.Add("d")

	intersection := l.Intersection(r)
	assert.Equal(t, 2, intersection.Count())
	assert.True(t, intersection.Is("b", "c"))

	// Shouldn't mutate originals
	assert.True(t, l.Is("a", "b", "c"))
	assert.True(t, r.Is("b", "c", "d"))
}

func TestSet_IntersectionAllInCommon_ResultIsAll(t *testing.T) {
	l := NewStringSet()
	l.Add("a")
	l.Add("b")
	l.Add("c")

	r := NewStringSet()
	r.Add("c")
	r.Add("a")
	r.Add("b")

	intersection := l.Intersection(r)
	assert.Equal(t, 3, intersection.Count())
	assert.True(t, intersection.Is("b", "c", "a"))

	// Shouldn't mutate originals
	assert.True(t, l.Is("a", "b", "c"))
	assert.True(t, r.Is("c", "a", "b"))
}

func TestSet_DifferenceAllInCommon_ResultIsNone(t *testing.T) {
	l := NewStringSet()
	l.Add("a")
	l.Add("b")
	l.Add("c")

	r := NewStringSet()
	r.Add("a")
	r.Add("c")
	r.Add("b")

	diff := l.Difference(r)
	assert.Equal(t, 0, diff.Count())
	assert.True(t, diff.Is())

	// Shouldn't mutate originals
	assert.True(t, l.Is("a", "b", "c"))
	assert.True(t, r.Is("b", "a", "c"))
}

func TestSet_DifferenceSomeInCommon_ResultIsNonCommon(t *testing.T) {
	l := NewStringSet()
	l.Add("a")
	l.Add("b")
	l.Add("c")

	r := NewStringSet()
	r.Add("b")
	r.Add("c")
	r.Add("d")

	diff := l.Difference(r)
	assert.Equal(t, 2, diff.Count())
	assert.True(t, diff.Is("a", "d"))

	// Shouldn't mutate originals
	assert.True(t, l.Is("a", "b", "c"))
	assert.True(t, r.Is("b", "c", "d"))
}

func TestSet_DifferenceNoneInCommon_ResultIsAll(t *testing.T) {
	l := NewStringSet()
	l.Add("a")
	l.Add("b")
	l.Add("c")

	r := NewStringSet()
	r.Add("x")
	r.Add("y")
	r.Add("z")

	diff := l.Difference(r)
	assert.Equal(t, 6, diff.Count())
	assert.True(t, diff.Is("a", "b", "c", "x", "y", "z"))

	// Shouldn't mutate originals
	assert.True(t, l.Is("a", "b", "c"))
	assert.True(t, r.Is("z", "y", "x"))
}

func TestSet_SubtractAllInCommon_ResultIsNone(t *testing.T) {
	l := NewStringSet()
	l.Add("a")
	l.Add("b")
	l.Add("c")

	r := NewStringSet()
	r.Add("a")
	r.Add("b")
	r.Add("c")

	leftSubRight := l.Subtract(r)
	assert.Equal(t, 0, leftSubRight.Count())
	assert.True(t, leftSubRight.Is())

	// Shouldn't mutate originals
	assert.True(t, l.Is("a", "b", "c"))
	assert.True(t, r.Is("a", "b", "c"))
}

func TestSet_SubtractSomeInCommon_ResultMissingCommon(t *testing.T) {
	l := NewStringSet()
	l.Add("a")
	l.Add("b")
	l.Add("c")

	r := NewStringSet()
	r.Add("b")
	r.Add("c")
	r.Add("d")

	leftSubRight := l.Subtract(r)
	assert.Equal(t, 1, leftSubRight.Count())
	assert.True(t, leftSubRight.Is("a"))

	// Shouldn't mutate originals
	assert.True(t, l.Is("a", "b", "c"))
	assert.True(t, r.Is("b", "c", "d"))
}

func TestSet_SubtractNoneInCommon_NoEffect(t *testing.T) {
	l := NewStringSet()
	l.Add("a")
	l.Add("b")
	l.Add("c")

	r := NewStringSet()
	r.Add("x")
	r.Add("y")
	r.Add("z")

	leftSubRight := l.Subtract(r)
	assert.Equal(t, 3, leftSubRight.Count())
	assert.True(t, leftSubRight.Is("a", "b", "c"))

	// Shouldn't mutate originals
	assert.True(t, l.Is("a", "b", "c"))
	assert.True(t, r.Is("x", "y", "z"))
}

func TestSet_Subtract_Noncommutative(t *testing.T) {
	l := NewStringSet()
	l.Add("a")
	l.Add("b")
	l.Add("c")

	r := NewStringSet()
	r.Add("b")
	r.Add("c")
	r.Add("d")

	leftSubRight := l.Subtract(r)
	assert.Equal(t, 1, leftSubRight.Count())
	assert.True(t, leftSubRight.Is("a"))

	rightSubLeft := r.Subtract(l)
	assert.Equal(t, 1, rightSubLeft.Count())
	assert.True(t, rightSubLeft.Is("d"))

	// Shouldn't mutate originals
	assert.True(t, l.Is("a", "b", "c"))
	assert.True(t, r.Is("b", "c", "d"))
}
