func pivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	sum := 0
	for i := range nums {
		if total - sum - nums[i] == sum {
			return i
		}
		sum += nums[i]
	}
	return -1
}
