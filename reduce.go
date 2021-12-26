package underscore

func Reduce[T1, T2 any](d []T1, h func(T1, T2) T2, init T2) T2 {
	val := init
	for _, v := range d {
		val = h(v, val)
	}
	return val
}
