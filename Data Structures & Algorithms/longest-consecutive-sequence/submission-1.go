func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})

	for _, value := range nums {
		m[value] = struct{}{}
	}

	result := 0
	for _, value := range nums {
		_, mid := m[value - 1]
		if mid {
			continue
		}
		l := 1
		for {
			_, ok := m[value + l]
			if ok {
				l++
			} else {
				break
			}
		}

		result = max(result, l)
	}

	return result
}
