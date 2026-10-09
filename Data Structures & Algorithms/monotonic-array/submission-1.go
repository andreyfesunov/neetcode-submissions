func isMonotonic(nums []int) bool {
    n := len(nums)
    increase := true
    for i := 1; i < n; i++ {
        if nums[i] < nums[i-1] {
            increase = false
            break
        }
    }
    if increase {
        return true
    }

    decrease := true
    for i := 1; i < n; i++ {
        if nums[i] > nums[i-1] {
            decrease = false
            break
        }
    }
    return decrease
}