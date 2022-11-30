package robot

import "fmt"

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

func (r *Step2Robot) Advance(extent Rect) {
	switch r.Dir {
	case N:
		if r.Northing < extent.Max.Northing {
			r.Northing++
		}
	case S:
		if r.Northing > extent.Min.Northing {
			r.Northing--
		}
	case E:
		if r.Easting < extent.Max.Easting {
			r.Easting++
		}
	case W:
		if r.Easting > extent.Min.Easting {
			r.Easting--
		}
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
