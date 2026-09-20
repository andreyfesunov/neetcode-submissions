func calPoints(operations []string) int {
	s := make([]int, 0)
	for _, v := range operations {
		i := len(s)
		switch v {
			case "+":
				s = append(s, s[i - 1] + s[i - 2])
			case "C":
				s = s[:i-1]
			case "D":
				s = append(s, s[i - 1] * 2)
			default:
				n, _ := strconv.Atoi(v)
				s = append(s, n)				
		}
	}
	return sum(s)
}

func sum(numbers []int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}