func scoreOfString(s string) int {
	x := []byte(s)
	res := 0

	for i := 1; i < len(x); i++ {
		res += abs(int(x[i]) - int(x[i - 1]))
	}

	return res
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}