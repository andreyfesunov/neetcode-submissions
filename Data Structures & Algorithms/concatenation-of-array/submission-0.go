func getConcatenation(nums []int) []int {
	result := make([]int, 2 * len(nums))

	offset := copy(result, nums)
	copy(result[offset:], nums)

	return result
}
