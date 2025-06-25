package slicex

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneral(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	assert.True(t, slices.Contains(items, 3))
	assert.False(t, slices.Contains(items, 6))

	v, ok := Search(items, func(item int) bool { return item == 3 })
	assert.True(t, ok)
	assert.Equal(t, 3, v)

	v, ok = Search(items, func(item int) bool { return item == 6 })
	assert.False(t, ok)
	assert.Equal(t, 0, v)
}

func TestUnique(t *testing.T) {
	assert.ElementsMatch(t, []string{"A", "B", "C"}, Unique([]string{"A", "B", "B", "B", "C", "C"}))
	assert.ElementsMatch(t, []int64{8, 7, 2, 3}, Unique([]int64{2, 2, 2, 2, 3, 3, 7, 7, 7, 7, 7, 8, 8, 8}))
}
