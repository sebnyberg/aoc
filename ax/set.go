package ax

type Set[T comparable] map[T]struct{}

func (s Set[T]) Has(k T) bool {
	_, exists := s[k]
	return exists
}

func (s Set[T]) Add(k T) {
	s[k] = struct{}{}
}
