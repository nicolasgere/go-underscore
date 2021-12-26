package main

func Find[T any](d []T, h func(T) bool) *T {

	for _, v := range d {
		if h(v) {
			return &v
		}
	}
	return nil
}
