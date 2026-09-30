func majorityElement(nums []int) int {
    c := make(map[int]int)
	m, v := 0, 0
	for i := range nums {
		c[nums[i]]++
		if c[nums[i]] > m {
			m = c[nums[i]]
			v = nums[i]
		}
	}
	return v
}
