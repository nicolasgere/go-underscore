package underscore

func Some[T any](d []T, h func(T) bool) bool {
	for _, v := range d {
		if h(v) {
			return true
		}
	}
	return false
}
