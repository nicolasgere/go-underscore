package underscore

func Every[T any](d []T, h func(T) bool) bool {
	for _, v := range d {
		if h(v) == false {
			return false
		}
	}
	return true
}
