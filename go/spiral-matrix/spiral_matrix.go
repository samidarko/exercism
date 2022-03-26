package spiralmatrix

func SpiralMatrix(size int) [][]int {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	x, y := 0, 0
	dx, dy := 1, 0
	for value := 1; value <= size*size; value++ {
		matrix[y][x] = value
		nextX, nextY := x+dx, y-dy
		if nextX >= size || nextY >= size || nextX < 0 || nextY < 0 || matrix[nextY][nextX] > 0 {
			dx, dy = dy, -dx
		}
		x += dx
		y -= dy
	}
	return matrix
}
