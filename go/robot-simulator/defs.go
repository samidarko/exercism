package robot

import (
	"errors"
	"fmt"
)

// definitions used in step 1

var Step1Robot struct {
	X, Y int
	Dir
}

type Dir int

var _ fmt.Stringer = Dir(1729)

// additional definitions used in step 2

type Command byte // valid values are 'R', 'L', 'A'
type RU int
type Pos struct{ Easting, Northing RU }
type Rect struct{ Min, Max Pos }
type Step2Robot struct {
	Dir
	Pos
}

var ErrOutsideRoom = errors.New("outside room")

func (r *Step2Robot) IsOutsideRoom(extent Rect) error {
	if r.Pos.Easting < extent.Min.Easting || r.Pos.Easting > extent.Max.Easting {
		return ErrOutsideRoom
	}
	if r.Pos.Northing < extent.Min.Northing || r.Pos.Northing > extent.Max.Northing {
		return ErrOutsideRoom
	}

	return nil
}

func (r *Step2Robot) Advance(extent Rect) error {

	outsideRoom := true

	switch r.Dir {
	case N:
		if r.Northing < extent.Max.Northing {
			r.Northing++
			outsideRoom = false
		}
	case S:
		if r.Northing > extent.Min.Northing {
			r.Northing--
			outsideRoom = false
		}
	case E:
		if r.Easting < extent.Max.Easting {
			r.Easting++
			outsideRoom = false
		}
	case W:
		if r.Easting > extent.Min.Easting {
			r.Easting--
			outsideRoom = false
		}
	}
	if outsideRoom {
		return ErrOutsideRoom
	}
	return nil
}

func (r *Step2Robot) Back() {
	switch r.Dir {
	case N:
		r.Northing--
	case S:
		r.Northing++
	case E:
		r.Easting--
	case W:
		r.Easting++
	}
}

func (r *Step2Robot) Right() {
	r.Dir = (r.Dir + 1) % 4
}

func (r *Step2Robot) Left() {
	r.Dir = (r.Dir + 3) % 4
}

func (r *Step2Robot) String() string {
	return fmt.Sprintf("Robot{%s, Pos{%d, %d}}", r.Dir, r.Pos.Easting, r.Pos.Northing)
}

// additional definition used in step 3

type Step3Robot struct {
	Name string
	Step2Robot
}
