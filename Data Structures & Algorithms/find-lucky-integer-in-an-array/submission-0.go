func findLucky(arr []int) int {
    sort.Ints(arr)
    streak := 0

    for i := len(arr) - 1; i >= 0; i-- {
        streak++
        if i == 0 || arr[i] != arr[i-1] {
            if arr[i] == streak {
                return arr[i]
            }
            streak = 0
        }
    }
    return -1
}