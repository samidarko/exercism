package stateoftictactoe

import (
	"errors"
)

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

type Line []rune

func (l Line) getState() (State, error) {
	if len(l) != 3 {
		return "", errors.New("line bad length")
	}

	for i := range l {
		if l[i] == ' ' {
			return Ongoing, nil
		}
		if l[0] != l[i] {
			return "", nil
		}
	}

	return Win, nil
}

type Lines []Line

func StateOfTicTacToe(board []string) (State, error) {
	xs, os := totalMoves(board)
	if !(xs == os+1 || xs == os) {
		return "", errors.New("incorrect moves")
	}
	lines := append(getDiagonalLines(board), append(getHorizontalLines(board), getVerticalLines(board)...)...)
	isOngoing, hasWon := false, false
	for _, line := range lines {
		state, err := line.getState()
		if err != nil {
			return "", err
		}
		if state == Win {
			if hasWon {
				return "", errors.New("players kept playing after a win")
			}
			hasWon = true
		}
		if state == Ongoing {
			isOngoing = true
		}
	}
	if hasWon {
		return Win, nil
	}
	if isOngoing {
		return Ongoing, nil
	}
	return Draw, nil
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

func totalMoves(board []string) (xs, os int) {
	for _, line := range board {
		for _, move := range line {
			if move == 'X' {
				xs++
			}
			if move == 'O' {
				os++
			}
		}
	}

	return
}
