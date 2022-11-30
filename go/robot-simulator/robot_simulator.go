package robot

import "fmt"

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

// Action type here.
type Action int

func StartRobot(command chan Command, action chan Action) {
	for c := range command {
		action <- Action(c)
	}

	close(action)
}

func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {

	for a := range action {
		switch a {
		case 'R':
			robot.Right()
		case 'L':
			robot.Left()
		case 'A':
			robot.Advance(extent)
		}
	}
	report <- robot
	close(report)
}

// Step 3

// Action3 type here.
type Action3 int

func StartRobot3(name, script string, action chan Action3, log chan string) {
	panic("Please implement the StartRobot3 function")
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {
	panic("Please implement the Room3 function")
}
