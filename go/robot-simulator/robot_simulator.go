package robot

import (
	"errors"
	"fmt"
)

// See defs.go for other definitions

// Step 1
// Define Dir type here.
// Define N, E, S, W here.
const (
	N Dir = iota
	E
	S
	W
)

func Right() {
	Step1Robot.Dir = (Step1Robot.Dir + 1) % 4
}

func Left() {
	Step1Robot.Dir = (Step1Robot.Dir - 1 + 4) % 4
}

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case S:
		Step1Robot.Y--
	case E:
		Step1Robot.X++
	case W:
		Step1Robot.X--
	}
}

func (d Dir) String() string {
	switch d {
	case N:
		return "N"
	case W:
		return "W"
	case S:
		return "S"
	case E:
		return "E"
	}
	panic(fmt.Sprintf("unknown direction %s", d.String()))
}

// Step 2

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

// Action type here.
type Action int

func StartRobot(command chan Command, action chan Action) {
	defer close(action)

	for c := range command {
		action <- Action(c)
	}
}

func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	for a := range action {
		switch a {
		case 'R':
			robot.Right()
		case 'L':
			robot.Left()
		case 'A':
			_ = robot.Advance(extent)
		}
	}
	report <- robot
}

// Step 3

// Action3 type here.
type Action3 struct {
	Name string
	Action
}

func StartRobot3(name, script string, action chan Action3, log chan string) {
	defer func() {
		action <- Action3{name, 'X'}
	}()
	if name == "" {
		log <- "no name"
		return
	}
	for _, c := range script {
		action <- Action3{name, Action(c)}
	}
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, report chan []Step3Robot, log chan string) {
	defer func() { report <- robots }()
	robotsMap := map[string]int{}

	for i, robot := range robots {
		if err := robot.IsOutsideRoom(extent); err != nil {
			log <- "outside room"
			return
		}
		if _, ok := robotsMap[robot.Name]; ok {
			log <- fmt.Sprint("duplicate name ", robot.Name)
			return
		}

		for _, robotIndex := range robotsMap {
			if robots[robotIndex].Pos == robot.Pos {
				log <- fmt.Sprintf("same position for %s and %s ", robot.Name, robots[robotIndex].Name)
				return
			}
		}

		robotsMap[robot.Name] = i
	}

	count := 0

	for a := range action {
		if _, ok := robotsMap[a.Name]; !ok {
			log <- fmt.Sprintf("bad robot %s ", a.Name)
			return
		}

		switch a.Action {
		case 'R':
			robots[robotsMap[a.Name]].Right()
		case 'L':
			robots[robotsMap[a.Name]].Left()
		case 'A':
			err := robots[robotsMap[a.Name]].Advance(extent)
			if err != nil {
				log <- "bump into wall"
			}

			for robotName, robotIndex := range robotsMap {
				if robots[robotsMap[a.Name]].Name != robotName && robots[robotsMap[a.Name]].Pos == robots[robotIndex].Pos {
					log <- "bump into each other"
					robots[robotsMap[a.Name]].Back()
				}
			}

		case 'X':
			count++
			if count == len(robots) {
				return
			}
		default:
			log <- fmt.Sprintf("bad command %c ", rune(a.Action))
			return
		}
	}
}
