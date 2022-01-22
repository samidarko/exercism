package pascal

// Triangle compute Pascal's triangle up to a given number of rows
func Triangle(n int) (triangle [][]int) {
	triangle = make([][]int, n)
	triangle[0] = []int{1}
	if n == 1 {
		return
	}
	for i := 1; i < n; i++ {
		previous := triangle[i-1]
		current := make([]int, i+1)
		current[0], current[i] = 1, 1
		for j := 1; j < i; j++ {
			current[j] = previous[j-1] + previous[j]
		}
		triangle[i] = current
	}
	return
}
