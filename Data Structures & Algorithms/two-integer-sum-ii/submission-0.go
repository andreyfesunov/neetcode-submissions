func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers) - 1
	for numbers[left] + numbers[right] != target {
		r := numbers[left] + numbers[right]
		if r > target {
			right--
		} else {
			left++
		}
	}
	return []int{left+1, right+1}
}
