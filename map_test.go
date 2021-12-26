package underscore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	c1 := Map([]int{1, 2, 3}, func(x int) int64 {
		return int64(x) + 10
	})
	assert.Equal(t, []int64{11, 12, 13}, c1)
}
