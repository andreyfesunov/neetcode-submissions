func findMaxConsecutiveOnes(nums []int) int {
	c, m := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			c++
		} else {
			c = 0
		}
		m = max(m, c)
	}
	return m
}
