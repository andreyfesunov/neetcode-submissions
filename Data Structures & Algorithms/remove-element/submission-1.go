func removeElement(nums []int, val int) int {
    tmp := []int{}
    for _, num := range nums {
        if num != val {
            tmp = append(tmp, num)
        }
    }
    for i := 0; i < len(tmp); i++ {
        nums[i] = tmp[i]
    }
    return len(tmp)
}