func sortedSquares(nums []int) []int {
	n := len(nums)
	l, r, id := 0, n - 1, n - 1 
	res := make([]int, n)

	for l <= r {
		if abs(nums[l]) > abs(nums[r]) {
			res[id] = nums[l] * nums[l]
			l++
		} else {
			res[id] = nums[r] * nums[r]
			r--
		}
		id--
	}

	return res
}

func abs(v int) int {
	if v < 0 {
		return - v
	}
	return v
}