package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReduce(t *testing.T) {
	c1 := Reduce([]int{1, 2, 3}, func(x int, val int) int {
		return val + x
	}, 0)
	assert.Equal(t, 6, c1)
}
