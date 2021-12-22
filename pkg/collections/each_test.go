package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEach(t *testing.T) {
	c1 := Each([]int{1, 2, 3}, func(x int) int {
		return x + 10
	})
	assert.Equal(t, []int{11, 12, 13}, c1)
}
