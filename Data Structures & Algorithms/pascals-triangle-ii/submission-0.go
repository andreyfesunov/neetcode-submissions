func getRow(rowIndex int) []int {
    res := []int{1}
    for i := 0; i < rowIndex; i++ {
        nextRow := make([]int, len(res)+1)
        for j := 0; j < len(res); j++ {
            nextRow[j] += res[j]
            nextRow[j+1] += res[j]
        }
        res = nextRow
    }
    return res
}