func productExceptSelf(nums []int) []int {
    r := 1
    zeros := make([]int, 0)
    for i, value := range nums {
        if value == 0 {
            zeros = append(zeros, i)
        } else {
            r *= value
        }
    }
    result := make([]int, len(nums))

    if len(zeros) == 1 {
        i := zeros[0]
        result[i] = r
    }
    if len(zeros) != 0 {
        return result
    }

    for i, value := range nums {
        result[i] = r / value
    }
    return result
}
