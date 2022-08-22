package robot

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
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = E
	case S:
		Step1Robot.Dir = W
	case E:
		Step1Robot.Dir = S
	case W:
		Step1Robot.Dir = N
	}
}

func Left() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = W
	case S:
		Step1Robot.Dir = E
	case E:
		Step1Robot.Dir = N
	case W:
		Step1Robot.Dir = S
	}
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
	panic("Please implement the String function")
}

// Step 2

// Action type here.
type Action int

func StartRobot(command chan Command, action chan Action) {
	panic("Please implement the StartRobot function")
}

func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	panic("Please implement the Room function")
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
