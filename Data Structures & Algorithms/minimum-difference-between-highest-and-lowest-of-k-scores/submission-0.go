func minimumDifference(nums []int, k int) int {
    sort.Ints(nums)
    l, r := 0, k-1
    res := math.MaxInt32

    for r < len(nums) {
        if nums[r]-nums[l] < res {
            res = nums[r] - nums[l]
        }
        l++
        r++
    }

    return res
}