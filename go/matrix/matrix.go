package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Matrix type
type Matrix struct {
	data []int
	rows int
	cols int
}

// New returns a new Matrix if no error
func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	var matrix Matrix
	matrix.rows = len(rows)
	for i, row := range rows {
		values := strings.Split(strings.TrimSpace(row), " ")
		if i == 0 {
			matrix.cols = len(values)
		}
		if i > 0 && len(values) != matrix.cols {
			return nil, fmt.Errorf("uneven rows")
		}
		for _, sValue := range values {
			value, err := strconv.Atoi(sValue)
			if err != nil {
				return nil, err
			}
			matrix.data = append(matrix.data, value)
		}
	}
	return &matrix, nil
}

// Cols must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	cols := make([][]int, m.cols)
	for c := 0; c < m.cols; c++ {
		col := make([]int, m.rows)
		for r := 0; r < m.rows; r++ {
			col[r] = m.data[c+r*m.cols]
		}
		cols[c] = col
	}
	return cols
}

// Rows must return the results without affecting the matrix.
func (m *Matrix) Rows() [][]int {
	rows := make([][]int, m.rows)
	for r := 0; r < m.rows; r++ {
		row := make([]int, m.cols)
		for c := 0; c < m.cols; c++ {
			row[c] = m.data[r*m.cols+c]
		}
		rows[r] = row
	}
	return rows
}

// Set must set a new value
func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 || row >= m.rows || col >= m.cols {
		return false
	}
	position := row*m.cols + col
	m.data[position] = val
	return true
}
