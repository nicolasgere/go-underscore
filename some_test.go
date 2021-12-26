package underscore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomeTrue(t *testing.T) {
	r := Some([]int{1, 2, 3}, func(x int) bool {
		return x == 2
	})
	assert.Equal(t, true, r)
}

func TestSomeFalse(t *testing.T) {
	r := Some([]int{1, 2, 3}, func(x int) bool {
		return x > 100
	})
	assert.Equal(t, false, r)
}
