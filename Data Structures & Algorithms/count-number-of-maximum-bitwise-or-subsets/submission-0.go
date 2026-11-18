func countMaxOrSubsets(nums []int) int {
    target := 0
    for _, x := range nums {
        target |= x
    }
    values := []int{0}
    for _, x := range nums {
        size := len(values)
        for i := 0; i < size; i++ {
            values = append(values, values[i]|x)
        }
    }
    answer := 0
    for _, v := range values[1:] {
        if v == target {
            answer++
        }
    }
    return answer
}
