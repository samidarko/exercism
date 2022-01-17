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
	for i := 0; i < m.cols; i++ {
		col := make([]int, m.rows)
		for j := 0; j < m.rows; j++ {
			col[j] = m.data[i+j*m.cols]
		}
		cols[i] = col
	}
	return cols
}

// Rows must return the results without affecting the matrix.
func (m *Matrix) Rows() [][]int {
	rows := make([][]int, m.rows)
	for i := 0; i < m.rows; i++ {
		row := make([]int, m.cols)
		for j := 0; j < m.cols; j++ {
			row[j] = m.data[i*m.cols+j]
		}
		rows[i] = row
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
