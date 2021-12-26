package underscore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvery(t *testing.T) {
	r := Every([]int{1, 2, 3}, func(x int) bool {
		return x < 100
	})
	assert.Equal(t, true, r)
}
