func validateStackSequences(pushed []int, popped []int) bool {
    stack := []int{}
    i := 0
    for _, n := range pushed {
        stack = append(stack, n)
        for i < len(popped) && len(stack) > 0 && popped[i] == stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
            i++
        }
    }
    return len(stack) == 0
}