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
				count += Explore(diagram, rowIndex, colIndex, Position{rowIndex: rowIndex, colIndex: colIndex}, None)
			}
		}
	}
	return count
}

func Explore(grid []string, rowIndex, colIndex int, initialPosition Position, direction Direction) int {
	rowInbounds := 0 <= rowIndex && rowIndex < len(grid)
	colInbounds := 0 <= colIndex && colIndex < len(grid[0])
	if !rowInbounds || !colInbounds {
		return 0
	}

	if direction != None && initialPosition.rowIndex == rowIndex && initialPosition.colIndex == colIndex {
		return 1
	}

	cell := grid[rowIndex][colIndex]

	//count := 0
	if cell == '+' && direction == None {
		//count := Explore(grid, rowIndex, colIndex+1, initialPosition, Right)
		//count += Explore(grid, rowIndex+1, colIndex, initialPosition, Down)
		return Explore(grid, rowIndex, colIndex+1, initialPosition, Right)
	}

	if cell == '-' && direction == Right {
		return Explore(grid, rowIndex, colIndex+1, initialPosition, Right)
	}

	if cell == '+' && direction == Right {
		return Explore(grid, rowIndex+1, colIndex, initialPosition, Down)
	}

	if cell == '|' && direction == Down {
		return Explore(grid, rowIndex+1, colIndex, initialPosition, Down)
	}

	if cell == '+' && direction == Down {
		return Explore(grid, rowIndex, colIndex-1, initialPosition, Left)
	}

	if cell == '-' && direction == Left {
		return Explore(grid, rowIndex, colIndex-1, initialPosition, Left)
	}

	if cell == '+' && direction == Left {
		return Explore(grid, rowIndex-1, colIndex, initialPosition, Up)
	}

	if cell == '|' && direction == Up {
		return Explore(grid, rowIndex-1, colIndex, initialPosition, Up)
	}

	return 0
}
