func search(nums []int, target int) int {
    left, right := 0, len(nums) - 1

    for right - left > 1 {
        mid := (left + right) / 2

        if nums[mid] > target {
            right = mid
        } else {
            left = mid
        }
    }

    if nums[left] == target {
        return left
    }
    if nums[right] == target {
        return right
    }

    return -1
}
