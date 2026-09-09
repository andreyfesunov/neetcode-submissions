func spiralOrder(matrix [][]int) []int {
    m, n := len(matrix), len(matrix[0])
	res := []int{}

	var dfs func(row, col, r, c, dr, dc int)
	dfs = func(row, col, r, c, dr, dc int) {
		if row == 0 || col == 0 {
			return
		}

		for i := 0; i < col; i++ {
			r += dr
			c += dc
			res = append(res, matrix[r][c])
		}

		dfs(col, row-1, r, c, dc, -dr)
	}

	dfs(m, n, 0, -1, 0, 1)
	return res
}
