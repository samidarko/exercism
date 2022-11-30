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
	defer close(action)

	for c := range command {
		action <- Action(c)
	}
}

func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	defer close(report)

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
}

// Step 3

// Action3 type here.
type Action3 struct {
	Name string
	Action
}

func StartRobot3(name, script string, action chan Action3, log chan string) {
	//if name == "" {
	//	log <- "no name"
	//	return
	//}
	for _, c := range script {
		action <- Action3{name, Action(c)}
	}
	//close(action)
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, report chan []Step3Robot, log chan string) {
	robotsMap := map[string]Step3Robot{}
	//defer close(action)
	//defer close(report)

	for _, robot := range robots {
		if robot.Name == "" {
			log <- "no name"
			close(report)
			return
		}
		if _, ok := robotsMap[robot.Name]; ok {
			log <- fmt.Sprint("duplicate name ", robot.Name)
			//close(log)
			close(report)
			return
		}

		for _, r := range robotsMap {
			if r.Pos == robot.Pos {
				log <- fmt.Sprintf("same position for %s and %s ", robot.Name, r.Name)
				close(report)
				return
			}
		}

		robotsMap[robot.Name] = robot
	}

	for a := range action {
		r := robotsMap[a.Name]
		switch a.Action {
		case 'R':
			r.Right()
		case 'L':
			r.Left()
		case 'A':
			r.Advance(extent)
		}
		robotsMap[a.Name] = r
	}
	report <- robots
	//close(report)
}
