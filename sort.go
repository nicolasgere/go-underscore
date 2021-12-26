package main

import (
	"math/rand"
)

func Sort[T any](d []T, h func(T, T) bool) []T {
	if len(d) < 2 {
		return d
	}

	left, right := 0, len(d)-1

	pivot := rand.Int() % len(d)

	d[pivot], d[right] = d[right], d[pivot]

	for i, _ := range d {
		if h(d[i], d[right]) {
			d[left], d[i] = d[i], d[left]
			left++
		}
	}

	d[left], d[right] = d[right], d[left]

	Sort(d[:left], h)
	Sort(d[left+1:], h)

	return d
}
