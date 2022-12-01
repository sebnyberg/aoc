package ax

type Stack[T any] []T

func (s *Stack[T]) Push(x T) {
	*s = append(*s, x)
}

func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		var nilT T
		return nilT
	}
	it := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return it
}

func (s *Stack[T]) Peek() T {
	if len(*s) == 0 {
		var nilT T
		return nilT
	}
	return (*s)[len(*s)-1]
}
