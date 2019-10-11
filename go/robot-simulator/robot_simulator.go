package robot

import (
	"fmt"
)

// N n
var N = Dir('N')

// W w
var W = Dir('W')

// S s
var S = Dir('S')

// E e
var E = Dir('E')

type Action byte
type Action3 int

func (d Dir) String() string {
	return ""
}

// Room implements the room logic
func Room(
	room Rect,
	robot Step2Robot,
	chanAction chan Action,
	chanStep2Robot chan Step2Robot) {
	// iterating: https://tour.golang.org/concurrency/5
	for action := range chanAction {
		fmt.Printf("ROOM# got action: %v (%s)\n", action, string(action))
		fmt.Printf("ROOM# current robot situation: %d %v\n",
			robot.Dir, robot.Pos)
		// robot will contain the data needed in order to move it.
		switch action {
		case 'A':
			fmt.Println("ROOM# advancing")
			robot.advance(room)
			break
		case 'L':
			fmt.Println("ROOM# left")
			robot.left()
			break
		case 'R':
			fmt.Println("ROOM# right")
			robot.right()
			break
		default:
			fmt.Println("ROOM# no valid action")
		}
		fmt.Printf("ROOM# temporary new robot situation: %d %v\n",
			robot.Dir, robot.Pos)
	}
	fmt.Printf("ROOM# new robot situation: %d %v\n",
		robot.Dir, robot.Pos)

	go func(chn chan Step2Robot) {
		fmt.Println("ROOM# sending robot via channel")
		chn <- robot
	}(chanStep2Robot)
}

func (r *Step2Robot) advance(room Rect) {
	switch r.Dir {
	case N:
		if r.Pos.Northing != room.Max.Northing {
			r.Pos.Northing++
		} else {
			fmt.Println("wall in N")
		}
		break
	case S:
		if r.Pos.Northing != room.Min.Northing {
			r.Pos.Northing--
		} else {
			fmt.Println("wall in S")
		}
		break
	case E:
		if r.Pos.Easting != room.Max.Easting {
			r.Pos.Easting++
		} else {
			fmt.Println("wall in E")
		}
		break
	case W:
		if r.Pos.Easting != room.Min.Easting {
			r.Pos.Easting--
		} else {
			fmt.Println("wall in W")
		}
		break
	}
}

func (r *Step2Robot) left() {
	switch r.Dir {
	case N:
		r.Dir = W
		break
	case S:
		r.Dir = E
		break
	case E:
		r.Dir = N
		break
	case W:
		r.Dir = S
		break
	}
}

func (r *Step2Robot) right() {
	switch r.Dir {
	case N:
		r.Dir = E
		break
	case S:
		r.Dir = W
		break
	case E:
		r.Dir = S
		break
	case W:
		r.Dir = N
		break
	}
}

func Room3(
	rect Rect,
	ar1 []Step3Robot,
	action3 chan Action3,
	chanStep chan []Step3Robot,
	cmd chan string) {
}

// StartRobot starts a robot
func StartRobot(chCommand chan Command, chAction chan Action) {
	fmt.Println("######################")
	fmt.Println("SR# starting new robot")

	for command := range chCommand {
		fmt.Printf("SR# received a command: %v (%s)\n", command, string(command))
		switch command {
		case 'A':
			fmt.Println("SR# advance")
			chAction <- 'A'
			break
		case 'R':
			fmt.Println("SR# turn right")
			chAction <- 'R'
			break
		case 'L':
			fmt.Println("SR# turn left")
			chAction <- 'L'
			break
		default:
			fmt.Println("SR# nothing to do")
			chAction <- ' '
		}
	}
	fmt.Println("SR# after range chCommand")
	// we need to close chan in order to inform Room that actions are done
	// This will end the range action loop and push to the robotChannel, the
	// updated robot position
	close(chAction)
}

func StartRobot3(
	str1 string,
	str2 string,
	action3 chan Action3,
	chn chan string) {
}

// Advance moves 1 step the Step1Robot in the direction
func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
		break
	case S:
		Step1Robot.Y--
		break
	case E:
		Step1Robot.X++
		break
	case W:
		Step1Robot.X--
		break
	}
}

// Right turns Step1Robot to the right
func Right() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = E
		break
	case S:
		Step1Robot.Dir = W
		break
	case E:
		Step1Robot.Dir = S
		break
	case W:
		Step1Robot.Dir = N
		break
	}
}

// Left turns Step1Robot to the left
func Left() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = W
		break
	case S:
		Step1Robot.Dir = E
		break
	case E:
		Step1Robot.Dir = N
		break
	case W:
		Step1Robot.Dir = S
		break
	}
}
