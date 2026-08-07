import (
	"errors"
	"slices"
)

type stack[T any] struct {
	s []T
}

func (s *stack[T]) Push(v T) {
	s.s = append(s.s, v)
}

func (s *stack[T]) Pop() (T, error) {
	l := len(s.s)
	
	if l == 0 {
		return *new(T), errors.New("stack is empty")
	}

	result := s.s[l - 1]
	s.s = s.s[:l - 1]

	return result, nil
}

func (s *stack[T]) IsEmpty() bool {
	return len(s.s) == 0
}

type brace byte

var opening = []brace{'{', '[', '('}

func (b brace) IsOpening() bool {
	return slices.Contains(opening, b)
}

var matching = map[brace]brace{
	'{': '}',
	'[': ']',
	'(': ')',
}

func (b brace) Matches(v brace) bool {
	if !b.IsOpening() {
		return false
	}

	return matching[b] == v
}

func isValid(input string) bool {
	s := new(stack[brace])
	
	for _, v := range []brace(input) {
		if v.IsOpening() {
			s.Push(v)
		} else {
			c, err := s.Pop()
			if err != nil {
				return false
			}

			ok := c.Matches(v)
			if !ok {
				return false
			}
		}
	}

	return s.IsEmpty()
}
