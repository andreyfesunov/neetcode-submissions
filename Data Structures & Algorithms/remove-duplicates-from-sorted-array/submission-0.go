func removeDuplicates(nums []int) int {
    seen := make(map[int]bool)
    var unique []int
    for _, num := range nums {
        if !seen[num] {
            seen[num] = true
            unique = append(unique, num)
        }
    }
    sort.Ints(unique)
    copy(nums, unique)
    return len(unique)
}