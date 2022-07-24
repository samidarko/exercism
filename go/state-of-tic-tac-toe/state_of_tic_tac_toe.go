package stateoftictactoe

import "fmt"

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

type Line []rune
type Lines []Line

func StateOfTicTacToe(board []string) (State, error) {
	for _, line := range getDiagonalLines(board) {
		fmt.Println(string(line))
	}
	return Win, nil
}

func getHorizontalLines(board []string) (lines Lines) {
	for i := range board {
		lines = append(lines, Line(board[i]))
	}
	return
}

func getVerticalLines(board []string) (lines Lines) {
	for i := range board {
		line := make(Line, 3)
		for j := range board {
			line[j] = rune(board[j][i])
		}

		lines = append(lines, line)
	}
	return
}

func getDiagonalLines(board []string) Lines {

	lines := make(Lines, 2)
	lines[0] = make(Line, 3)
	lines[1] = make(Line, 3)

	for i, j := 0, 2; i <= 2; i, j = i+1, j-1 {
		lines[0][i] = rune(board[i][i])
		lines[1][i] = rune(board[j][i])
	}

	return lines
}
