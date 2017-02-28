package set

// This assumes a concrete type (StringSet) has been generated.

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetNew(t *testing.T) {
	set := NewStringSet()

	assert.NotNil(t, set)
}

func TestSetNewWithInitial(t *testing.T) {
	set := NewStringSet("a", "b")

	assert.Equal(t, 2, set.Count())
	assert.True(t, set.Is("a", "b"))
}

func TestSetAdd(t *testing.T) {
	set := NewStringSet()

	set.Add("a")

	assert.Equal(t, 1, set.Count())
	assert.True(t, set.Is("a"))
}

func TestSetRemove(t *testing.T) {
	set := NewStringSet()

	set.Add("a")
	set.Add("b")
	assert.Equal(t, 2, set.Count())

	set.Remove("b")
	assert.Equal(t, 1, set.Count())
}

func TestSetContains(t *testing.T) {
	set := NewStringSet()

	set.Add("a")

	assert.True(t, set.Contains("a"))
	assert.False(t, set.Contains("b"))
	assert.False(t, set.Contains(""))
}

func TestSetIs(t *testing.T) {
	set := NewStringSet()

	assert.True(t, set.Is())
	set.Add("a")
	assert.True(t, set.Is("a"))
	set.Add("b")
	set.Add("c")
	assert.True(t, set.Is("a", "b", "c"))

	assert.False(t, set.Is("a", "b"))
	assert.False(t, set.Is("a", "b", "x"))
	assert.False(t, set.Is("x", "b", "c"))
}

func TestSetCount(t *testing.T) {
	set := NewStringSet()

	assert.Zero(t, set.Count())

	set.Add("a")
	assert.Equal(t, 1, set.Count())

	set.Add("b")
	set.Add("c")
	assert.Equal(t, 3, set.Count())
}

func TestSetElements(t *testing.T) {
	set := NewStringSet()

	set.Add("a")
	assert.True(t, set.Is("a"))

	set.Add("b")
	set.Add("c")
	assert.True(t, set.Is("a", "b", "c"))

	assert.Equal(t, set.Count(), len(set.Elements()))
}

func TestSetClear(t *testing.T) {
	set := NewStringSet()

	set.Add("a")
	set.Add("b")
	set.Add("c")

	set.Clear()
	assert.Zero(t, set.Count())
}

func TestSetUnion(t *testing.T) {
	l := NewStringSet()
	l.Add("a")
	l.Add("b")
	l.Add("c")

	r := NewStringSet()
	r.Add("b")
	r.Add("c")
	r.Add("d")

	union := l.Union(r)
	assert.Equal(t, 4, union.Count())
	assert.True(t, union.Is("a", "b", "c", "d"))

	// Shouldn't mutate originals
	assert.True(t, l.Is("a", "b", "c"))
	assert.True(t, r.Is("b", "c", "d"))
}

func TestSetIntersection(t *testing.T) {
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

func TestSetDifference(t *testing.T) {
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

func TestSetSubtract(t *testing.T) {
	l := NewStringSet()
	l.Add("a")
	l.Add("b")
	l.Add("c")

	r := NewStringSet()
	r.Add("b")
	r.Add("c")
	r.Add("d")

	// Subtraction is noncommutative.

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
