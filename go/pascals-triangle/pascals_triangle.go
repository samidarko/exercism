package pascal

// Triangle compute Pascal's triangle up to a given number of rows
func Triangle(n int) [][]int {
	triangle := make([][]int, n)
	for i := 0; i < n; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][i] = 1, 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	return triangle
}
