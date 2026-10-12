func numIdenticalPairs(nums []int) int {
    count := make(map[int]int)
    res := 0
    for _, num := range nums {
        count[num]++
    }
    for _, c := range count {
        res += c * (c - 1) / 2
    }
    return res
}