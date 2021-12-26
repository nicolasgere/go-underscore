package main

func Each[T any](d []T, h func(T) T) []T {
	r := make([]T, len(d))
	for i, v := range d {
		r[i] = h(v)
	}
	return r
}
