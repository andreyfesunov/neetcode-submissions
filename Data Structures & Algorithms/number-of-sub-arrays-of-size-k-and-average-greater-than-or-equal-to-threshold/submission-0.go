func numOfSubarrays(arr []int, k int, threshold int) int {
    res := 0
    l := 0

    for r := k - 1; r < len(arr); r++ {
        sum := 0
        for i := l; i <= r; i++ {
            sum += arr[i]
        }
        if sum/k >= threshold {
            res++
        }
        l++
    }
    return res
}