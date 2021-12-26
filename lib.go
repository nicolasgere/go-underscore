package underscore

type Collection[T any] struct {
	Data []T
}

func New[T any](data []T) Collection[T] {
	c := Collection[T]{
		Data: data,
	}
	return c
}

func (self *Collection[T]) Filter(handler func(T) bool) Collection[T] {
	tmp := []T{}
	for _, v := range self.Data {
		if handler(v) {
			tmp = append(tmp, v)
		}
	}
	return New[T](tmp)
}
