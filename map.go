package underscore

func Map[T1, T2 any](d []T1, h func(T1) T2) []T2 {
	r := make([]T2, len(d))
	for i, v := range d {
		r[i] = h(v)
	}
	return r
}
