package rectangles

type Position struct {
	rowIndex int
	colIndex int
}

func NewPosition(rowIndex, colIndex int) Position {
	return Position{rowIndex: rowIndex, colIndex: colIndex}
}

func (p Position) Right() Position {
	return NewPosition(p.rowIndex, p.colIndex+1)
}

func (p Position) Left() Position {
	return NewPosition(p.rowIndex, p.colIndex-1)
}
func (p Position) Up() Position {
	return NewPosition(p.rowIndex-1, p.colIndex)
}

func (p Position) Down() Position {
	return NewPosition(p.rowIndex+1, p.colIndex)
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
	Idle
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
				count += Explore(diagram, NewPosition(rowIndex, colIndex), NewPosition(rowIndex, colIndex), Idle)
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

	if direction != Idle && currentPosition == initialPosition {
		return 1
	}

	cell := grid[currentPosition.rowIndex][currentPosition.colIndex]

	if cell == '+' && direction == Right {
		return Explore(grid, currentPosition.Down(), initialPosition, Down) + Explore(grid, currentPosition.Right(), initialPosition, Right)
	}

	if cell == '+' && direction == Left {
		return Explore(grid, currentPosition.Up(), initialPosition, Up) + Explore(grid, currentPosition.Left(), initialPosition, Left)
	}

	if cell == '+' && direction == Down {
		return Explore(grid, currentPosition.Left(), initialPosition, Left) + Explore(grid, currentPosition.Down(), initialPosition, Down)
	}

	if direction == Up && (cell == '|' || cell == '+') {
		return Explore(grid, currentPosition.Up(), initialPosition, Up)
	}

	if (cell == '-' && direction == Right) || (cell == '+' && direction == Idle) {
		return Explore(grid, currentPosition.Right(), initialPosition, Right)
	}

	if cell == '-' && direction == Left {
		return Explore(grid, currentPosition.Left(), initialPosition, Left)
	}

	if cell == '|' && direction == Down {
		return Explore(grid, currentPosition.Down(), initialPosition, Down)
	}

	return 0
}
