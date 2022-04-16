// Package bowling score a bowling game
package bowling

import (
	"fmt"
)

// Tabulation type
type Tabulation int

const (
	OpenFrame Tabulation = iota
	Spare
	Strike
)

// Frame type
type Frame struct {
	throws     []int
	tabulation Tabulation
}

// TotalPinsCount total pins count
func (f *Frame) TotalPinsCount() (score int) {
	for _, v := range f.throws {
		score += v
	}
	return
}

// ThrowsCount returns throws count of a frame
func (f *Frame) ThrowsCount() int {
	return len(f.throws)
}

// Game type here.
type Game struct {
	frames  [10]Frame
	current int // current game
}

// CurrentFrame returns the current Frame
func (g *Game) CurrentFrame() *Frame {
	return &g.frames[g.current]
}

// NewGame returns a new Game
func NewGame() *Game {
	return &Game{}
}

// Roll add the number of pins down for a throw
func (g *Game) Roll(pins int) error {
	if pins < 0 || pins > 10 {
		return fmt.Errorf("throws should be between 0 and 10 but it was %d", pins)
	}

	if g.current > 9 {
		return fmt.Errorf("game is over")
	}

	g.CurrentFrame().throws = append(g.CurrentFrame().throws, pins)

	switch {
	case g.CurrentFrame().TotalPinsCount() == 10 && g.CurrentFrame().ThrowsCount() == 1:
		g.CurrentFrame().tabulation = Strike
	case g.CurrentFrame().TotalPinsCount() == 10 && g.CurrentFrame().ThrowsCount() == 2:
		g.CurrentFrame().tabulation = Spare
	}

	// validate score
	if g.current < 9 && g.CurrentFrame().TotalPinsCount() > 10 {
		return fmt.Errorf("cannot score more than 10 points")
	}
	if g.current > 8 && g.CurrentFrame().tabulation == Spare && g.CurrentFrame().TotalPinsCount() > 20 {
		return fmt.Errorf("cannot score more than 20 points")
	}
	if g.current > 8 && g.CurrentFrame().tabulation == Strike && g.CurrentFrame().ThrowsCount() == 3 {
		if g.CurrentFrame().throws[1] < 10 && g.CurrentFrame().TotalPinsCount() > 20 {
			return fmt.Errorf("cannot score more than 20 points")
		}
		if g.CurrentFrame().throws[1] == 10 && g.CurrentFrame().TotalPinsCount() > 30 {
			return fmt.Errorf("cannot score more than 30 points")
		}
	}

	if (g.CurrentFrame().tabulation == OpenFrame && g.CurrentFrame().ThrowsCount() == 2) ||
		((g.CurrentFrame().tabulation == Strike || g.CurrentFrame().tabulation == Spare) && g.current < 9) ||
		g.CurrentFrame().ThrowsCount() == 3 { // can reach 3 throws on last frame
		g.current++
	}

	return nil
}

// FrameScore returns the score for a given frame index
func (g *Game) FrameScore(i int) int {
	if i < 0 || i > 9 {
		return 0
	}

	if i == 9 {
		return g.frames[i].TotalPinsCount()
	}

	score := g.frames[i].TotalPinsCount()
	switch g.frames[i].tabulation {
	case Strike:
		if len(g.frames[i+1].throws) > 1 {
			score += g.frames[i+1].throws[0] + g.frames[i+1].throws[1]
		} else {
			score += g.frames[i+1].throws[0] + g.frames[i+2].throws[0]
		}
	case Spare:
		score += g.frames[i+1].throws[0]
	}

	return score
}

// Score returns the score
func (g *Game) Score() (int, error) {
	if ((g.frames[9].tabulation == Strike || g.frames[9].tabulation == Spare) && len(g.frames[9].throws) < 3) || // last frame bonus roll not rolled
		g.current < 9 {
		return 0, fmt.Errorf("incomplete game")
	}
	score := 0
	for i := 0; i <= 9; i++ {
		score += g.FrameScore(i)
	}
	return score, nil
}
