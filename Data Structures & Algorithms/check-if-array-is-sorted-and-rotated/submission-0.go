func check(nums []int) bool {
    count, N := 0, len(nums)

    for i := 0; i < N; i++ {
        if nums[i] > nums[(i+1)%N] {
            count++
            if count > 1 {
                return false
            }
        }
    }

    return true
}