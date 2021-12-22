package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFind(t *testing.T) {
	c1 := Find([]int{1, 2, 3}, func(x int) bool {
		return x > 2
	})
	if c1 != nil {
		assert.Equal(t, 3, *c1)
	}
}
