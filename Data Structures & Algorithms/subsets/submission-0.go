func subsets(nums []int) [][]int {
    n := len(nums)
    res := [][]int{}

    for i := 0; i < (1 << n); i++ {
        subset := []int{}
        for j := 0; j < n; j++ {
            if (i & (1 << j)) != 0 {
                subset = append(subset, nums[j])
            }
        }
        res = append(res, subset)
    }

    return res
}
