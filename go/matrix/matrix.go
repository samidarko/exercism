package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Matrix type
type Matrix [][]int

// New returns a new Matrix if no error
func New(input string) (Matrix, error) {
	split := strings.Split(input, "\n")
	var matrix Matrix
	n := 0
	for i, s := range split {
		values := strings.Split(strings.TrimSpace(s), " ")
		if i == 0 {
			n = len(values)
		}
		if i > 0 && len(values) != n {
			return nil, fmt.Errorf("bad number of columns")
		}
		row := make([]int, n)
		for j, sValue := range values {
			value, err := strconv.Atoi(sValue)
			if err != nil {
				return nil, err
			}
			row[j] = value
		}
		matrix = append(matrix, row)
	}
	return matrix, nil
}

// Cols must return the results without affecting the matrix.
func (matrix Matrix) Cols() [][]int {
	m := len(matrix)
	n := len(matrix[0])
	cols := make([][]int, n)
	for c := 0; c < n; c++ {
		col := make([]int, m)
		for r := 0; r < m; r++ {
			col[r] = matrix[r][c]
		}
		cols[c] = col
	}
	return cols
}

// Rows must return the results without affecting the matrix.
func (matrix Matrix) Rows() [][]int {
	m := len(matrix)
	n := len(matrix[0])
	rows := make([][]int, m)
	for r := 0; r < m; r++ {
		row := make([]int, n)
		for c := 0; c < n; c++ {
			row[c] = matrix[r][c]
		}
		rows[r] = row
	}
	return rows
}

// Set must set a new value
func (matrix Matrix) Set(row, col, val int) bool {
	m := len(matrix)
	n := len(matrix[0])
	if row < 0 || col < 0 || row >= m || col >= n {
		return false
	}
	matrix[row][col] = val
	return true
}
