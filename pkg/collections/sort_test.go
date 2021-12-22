package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	r := Sort([]int{34, 10, 2, 50}, func(x int, y int) bool {
		return x < y
	})
	assert.Equal(t, []int{2, 10, 34, 50}, r)
}
