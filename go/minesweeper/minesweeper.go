package minesweeper

type Position struct {
	rowIndex int
	colIndex int
}

// Annotate returns an annotated board
func Annotate(board []string) []string {
	editableBoard := make([][]uint8, len(board))

	for rowIndex := range board {
		for colIndex := range board[rowIndex] {
			char := board[rowIndex][colIndex]
			if char == ' ' {
				var minesCount uint8 = '0'

				for _, position := range getSurroundingPositions(board, rowIndex, colIndex) {
					if board[position.rowIndex][position.colIndex] == '*' {
						minesCount++
					}
				}

				if minesCount > '0' {
					editableBoard[rowIndex] = append(editableBoard[rowIndex], minesCount)
				} else {
					editableBoard[rowIndex] = append(editableBoard[rowIndex], ' ')
				}

			}

			if char == '*' {
				editableBoard[rowIndex] = append(editableBoard[rowIndex], char)
			}
		}
	}

	result := make([]string, 0)

	for _, row := range editableBoard {
		result = append(result, string(row))
	}

	return result
}

func getSurroundingPositions(board []string, rowIndex, colIndex int) (positions []Position) {
	// (rowIndex-1, colIndex-1)(rowIndex-1, colIndex)(rowIndex-1, colIndex+1)
	// (rowIndex,   colIndex-1)(rowIndex,   colIndex)(rowIndex,   colIndex+1)
	// (rowIndex+1, colIndex-1)(rowIndex+1, colIndex)(rowIndex+1, colIndex+1)

	moves := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for _, move := range moves {
		position := Position{rowIndex: rowIndex + move[0], colIndex: colIndex + move[1]}

		if position.rowIndex >= 0 && position.rowIndex < len(board) && position.colIndex >= 0 && position.colIndex < len(board[0]) {
			positions = append(positions, position)
		}
	}

	return
}
