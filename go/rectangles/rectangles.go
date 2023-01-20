package rectangles

type Position struct {
	colIndex int
	rowIndex int
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
	None
)

func Count(diagram []string) int {
	if len(diagram) == 0 || len(diagram[0]) == 0 {
		return 0
	}
	rowsCount := len(diagram)
	colsCount := len(diagram[0]) // assumed grid will never be of length 0
	count := 0
	for rowIndex := 0; rowIndex < rowsCount; rowIndex++ {
		for colIndex := 0; colIndex < colsCount; colIndex++ {
			if diagram[rowIndex][colIndex] == '+' {
				count += Explore(diagram, Position{rowIndex: rowIndex, colIndex: colIndex}, Position{rowIndex: rowIndex, colIndex: colIndex}, None)
			}
		}
	}
	return count
}

func Explore(grid []string, currentPosition, initialPosition Position, direction Direction) int {
	rowInbounds := 0 <= currentPosition.rowIndex && currentPosition.rowIndex < len(grid)
	colInbounds := 0 <= currentPosition.colIndex && currentPosition.colIndex < len(grid[0])
	if !rowInbounds || !colInbounds {
		return 0
	}

	if direction != None && initialPosition.rowIndex == currentPosition.rowIndex && initialPosition.colIndex == currentPosition.colIndex {
		return 1
	}

	cell := grid[currentPosition.rowIndex][currentPosition.colIndex]

	if cell == '+' && direction == None {
		currentPosition.colIndex++
		return Explore(grid, currentPosition, initialPosition, Right)
	}

	if cell == '+' && direction == Right {
		currentPosition.rowIndex++
		return Explore(grid, currentPosition, initialPosition, Down)
	}

	if cell == '+' && direction == Left {
		currentPosition.rowIndex--
		return Explore(grid, currentPosition, initialPosition, Up)
	}

	if cell == '+' && direction == Down {
		currentPosition.colIndex--
		return Explore(grid, currentPosition, initialPosition, Left)
	}

	if cell == '-' && direction == Right {
		currentPosition.colIndex++
		return Explore(grid, currentPosition, initialPosition, Right)
	}

	if cell == '-' && direction == Left {
		currentPosition.colIndex--
		return Explore(grid, currentPosition, initialPosition, Left)
	}

	if cell == '|' && direction == Down {
		currentPosition.rowIndex++
		return Explore(grid, currentPosition, initialPosition, Down)
	}

	if cell == '|' && direction == Up {
		currentPosition.rowIndex--
		return Explore(grid, currentPosition, initialPosition, Up)
	}

	return 0
}
