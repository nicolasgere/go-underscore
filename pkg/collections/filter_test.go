package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	c1 := Filter([]int{1, 2, 3}, func(x int) bool {
		return x > 2
	})
	assert.Equal(t, []int{3}, c1)
}
