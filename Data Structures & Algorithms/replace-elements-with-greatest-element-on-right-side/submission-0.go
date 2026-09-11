func replaceElements(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	m := -1

	for i := n - 1; i >= 0; i-- {
		result[i] = m
		m = max(m, arr[i])
	}

	return result
}
