package underscore

func Filter[T any](d []T, h func(T) bool) []T {
	r := []T{}
	for _, v := range d {
		if h(v) {
			r = append(r, v)
		}
	}
	return r
}
