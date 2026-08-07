type stack struct {
	s []int
}

func (s *stack) Push(v int) {
	s.s = append(s.s, v)
}

func (s *stack) Pop() int {
	l := len(s.s)
	r := s.s[l - 1]
	s.s = s.s[:l - 1]
	return r
}

func evalRPN(tokens []string) int {
	if len(tokens) == 1 {
		result, _ := strconv.Atoi(tokens[0])
		return result
	}

	stack := new(stack)

	for _, token := range tokens {
		switch {
			case token == "+":
				stack.Push(stack.Pop() + stack.Pop())
			case token == "-":
				a, b := stack.Pop(), stack.Pop()
				stack.Push(b - a)
			case token == "*":
				stack.Push(stack.Pop() * stack.Pop())
			case token == "/":
				a, b := stack.Pop(), stack.Pop()
				stack.Push(b / a)
			default:
				entry, _ := strconv.Atoi(token)
				stack.Push(entry)
		}
	}

	return stack.Pop()
}
