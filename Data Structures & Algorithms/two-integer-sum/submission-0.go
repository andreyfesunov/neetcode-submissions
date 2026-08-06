func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

	for key, value := range nums {
		m[value] = key
	}

	for key1, value := range nums {
		key2, ok := m[target - value]

		if ok && key1 != key2 {
			return []int{key1, key2}
		}
	}

	return []int{}
}
