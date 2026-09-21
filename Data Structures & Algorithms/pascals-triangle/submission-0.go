func generate(numRows int) [][]int {
	triangle := make([][]int, 0, numRows)
	for i := 1; i <= numRows; i++ {
		triangle = append(triangle, make([]int, i))
	}
	triangle[0][0] = 1
	for row := 1; row < numRows; row++ {
		for i := 0; i <= row; i++ {
			if i == 0 || i == row {
				triangle[row][i] = 1
			} else {
				triangle[row][i] = triangle[row - 1][i - 1] + triangle[row - 1][i]
			}
		}
	}
	return triangle
}
