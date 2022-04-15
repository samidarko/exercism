package bowling

import (
	"fmt"
)

type Tabulation int

const (
	OpenFrame Tabulation = iota
	Spare
	Strike
)

type Frame struct {
	throws     []int
	tabulation Tabulation
}

func (f *Frame) Score() (score int) {
	for _, v := range f.throws {
		score += v
	}
	return
}

func (f *Frame) ThrowsCount() int {
	return len(f.throws)
}

// Game type here.
type Game struct {
	frames  [10]Frame
	current int // current game
}

func (g *Game) CurrentFrame() *Frame {
	return &g.frames[g.current]
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Roll(pins int) error {
	if pins < 0 || pins > 10 {
		return fmt.Errorf("throws should be between 0 and 10 but it was %d", pins)
	}

	if g.current > 9 {
		return fmt.Errorf("game is over")
	}

	g.CurrentFrame().throws = append(g.CurrentFrame().throws, pins)

	switch {
	case g.CurrentFrame().Score() == 10 && g.CurrentFrame().ThrowsCount() == 1:
		g.CurrentFrame().tabulation = Strike
	case g.CurrentFrame().Score() == 10 && g.CurrentFrame().ThrowsCount() == 2:
		g.CurrentFrame().tabulation = Spare
	}

	// validate score
	switch {
	case g.CurrentFrame().Score() > 10 && g.current < 9:
		return fmt.Errorf("cannot score more than 10 points")
	case g.CurrentFrame().Score() > 20 && g.current > 8 && (g.CurrentFrame().tabulation == Strike || g.CurrentFrame().tabulation == Spare): // last frame score validation
		return fmt.Errorf("cannot score more than 20 points")
	}

	if (g.CurrentFrame().tabulation == OpenFrame && g.CurrentFrame().ThrowsCount() == 2) ||
		((g.CurrentFrame().tabulation == Strike || g.CurrentFrame().tabulation == Spare) && g.current < 9) ||
		g.CurrentFrame().ThrowsCount() == 3 { // can reach 3 throws on last frame
		g.current++
	}

	return nil
}

func (g *Game) Score() (int, error) {
	panic("Please implement the Score function")
}
