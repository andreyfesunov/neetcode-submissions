func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])
    l, r := 0, m * n - 1

    for l < r {
        mid := l + (r - l) / 2
        mmid, nmid := mid / n, mid % n

        if matrix[mmid][nmid] > target {
            r = mid - 1
        } else if matrix[mmid][nmid] < target {
            l = mid + 1
        } else {
            return true
        }
    } 

    if matrix[l / n][l % n] == target {
        return true
    }

    return false
}
