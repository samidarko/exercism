package matrix

import (
	"strconv"
	"strings"
)

// Pair type
type Pair [2]int

// Matrix type
type Matrix struct {
	data [][]int
}

// New returns a matrix from string
func New(s string) (*Matrix, error) {
	matrix := new(Matrix)
	if s == "" {
		return matrix, nil
	}

	for _, row := range strings.Split(s, "\n") {
		data := make([]int, 0)
		for _, element := range strings.Split(row, " ") {
			value, err := strconv.Atoi(element)
			if err != nil {
				return nil, err
			}
			data = append(data, value)
		}

		matrix.data = append(matrix.data, data)
	}
	return matrix, nil
}

// Pairs returns all the matrix's pairs
func (m *Matrix) Pairs() (pairs []Pair) {
	// row / colum
	for rowNumber, row := range m.data {
		for colNumber := range row {
			pairs = append(pairs, [2]int{rowNumber + 1, colNumber + 1})
		}
	}
	return
}

// IsSaddle returns true if pair is a saddle
func (m *Matrix) IsSaddle(pair Pair) bool {
	return m.IsGreaterInRow(pair) && m.IsSmallerInCol(pair)
}

// IsGreaterInRow returns true if pair is greater or equal in its row
func (m *Matrix) IsGreaterInRow(pair Pair) bool {
	row, colum := pair[0], pair[1]
	pairValue := m.data[row-1][colum-1]

	for _, cellValue := range m.data[row-1] {
		if cellValue > pairValue {
			return false
		}
	}

	return true
}

// IsSmallerInCol returns true if pair is smallest or equal in its col
func (m *Matrix) IsSmallerInCol(pair Pair) bool {
	_, colum := pair[0], pair[1]
	pairValue := m.data[pair[0]-1][colum-1]

	for row := range m.data {
		cellValue := m.data[row][colum-1]
		if cellValue < pairValue {
			return false
		}
	}

	return true
}

// Saddle returns a list of saddles
func (m *Matrix) Saddle() (saddles []Pair) {
	for _, pair := range m.Pairs() {
		if m.IsSaddle(pair) {
			saddles = append(saddles, pair)
		}
	}
	return
}
