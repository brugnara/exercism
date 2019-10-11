package robot

import (
	"fmt"
	"strings"
)

// N n
var N = Dir('N')

// W w
var W = Dir('W')

// S s
var S = Dir('S')

// E e
var E = Dir('E')

// Action type
type Action byte

// Action3 type
type Action3 struct {
	robotName string
	action    Action
}

// junk, just to avoid build errors/warnings
func (d Dir) String() string {
	return ""
}

// Step 1

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

// Step 2

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

	fmt.Println("ROOM# sending robot via channel")
	chanStep2Robot <- robot
}

func (r *Step2Robot) advance(room Rect) (moved bool) {
	fmt.Println("we are moving!")
	switch r.Dir {
	case N:
		if r.Pos.Northing != room.Max.Northing {
			r.Pos.Northing++
			fmt.Println("/\\")
			moved = true
		} else {
			fmt.Println("wall in N")
		}
		break
	case S:
		if r.Pos.Northing != room.Min.Northing {
			r.Pos.Northing--
			fmt.Println("\\/")
			moved = true
		} else {
			fmt.Println("wall in S")
		}
		break
	case E:
		if r.Pos.Easting != room.Max.Easting {
			r.Pos.Easting++
			fmt.Println(">")
			moved = true
		} else {
			fmt.Println("wall in E")
		}
		break
	case W:
		if r.Pos.Easting != room.Min.Easting {
			r.Pos.Easting--
			fmt.Println("<")
			moved = true
		} else {
			fmt.Println("wall in W")
		}
		break
	default:
		fmt.Println("No valid direction provided :(")
	}
	fmt.Printf("has moved? %v\n", moved)
	return
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

// Step 3

// Room3 implementation
func Room3(
	room Rect,
	robots []Step3Robot,
	chanActions chan Action3,
	chanRobot chan []Step3Robot,
	log chan string) {
	//
	robotCount := len(robots)
	fmt.Printf("R3# received %d robot/s:\n", robotCount)

	names := make(map[string]bool)
	for _, robot := range robots {
		// dups spotter
		if names[robot.Name] {
			// dup!
			log <- "Duplicated robot name found :("
			continue
		}
		names[robot.Name] = true

		// checking initial robot positions
		msgOut := "out of borders"
		placeable := true
		if robot.Pos.Easting < room.Min.Easting || robot.Pos.Easting > room.Max.Easting {
			log <- msgOut
			placeable = false
		}
		if robot.Pos.Northing < room.Min.Northing || robot.Pos.Northing > room.Max.Northing {
			log <- msgOut
			placeable = false
		}

		if placeable {
			fmt.Println("Robot is placeable")
		}

	}

	for action := range chanActions {
		fmt.Printf("R3# got action: %v\n", action)

		if action.action == 'Q' {
			fmt.Println("R3# got Quit command")
			break
		}

		// search the robot we want to operate
		aRobotIsFound := false

		for index, robot := range robots {
			if robot.Name == action.robotName {
				aRobotIsFound = true
				fmt.Println("R3# operating robot:", action.robotName)
				fmt.Printf("R3# current robot situation: %d %v\n",
					robots[index].Step2Robot.Dir, robots[index].Step2Robot.Pos)
				switch action.action {
				case 'A':
					fmt.Println("R3# advancing")
					hasMoved := robots[index].Step2Robot.advance(room)
					if !hasMoved {
						log <- "robot over a wall"
					}
					break
				case 'L':
					fmt.Println("R3# left")
					robots[index].Step2Robot.left()
					break
				case 'R':
					fmt.Println("R3# right")
					robots[index].Step2Robot.right()
					break
				default:
					fmt.Println("R3# no valid action")
				}
				fmt.Printf("R3# new robot situation: %d %v\n",
					robots[index].Step2Robot.Dir, robots[index].Step2Robot.Pos)
			}
		}

		if !aRobotIsFound {
			msg := "No robot found with name: " + action.robotName
			fmt.Println(msg)
			log <- msg
		}

	}

	// finally
	fmt.Println("R3# sending updated robots")
	chanRobot <- robots
}

// StartRobot3 implementation
func StartRobot3(
	robotName string,
	script string,
	action chan Action3,
	log chan string) {
	fmt.Println("##################")
	fmt.Printf("SR3# robotName: %s\n", robotName)
	fmt.Println("SR3# got script:")
	fmt.Println(script)

	// defer close(action)
	defer func() {
		fmt.Println("Sending final 'action'")
		action <- Action3{
			action: 'Q',
		}
	}()

	if robotName == "" {
		fmt.Println("SR3# missing robot name")
		log <- "Missing robotName"
		return
	}

	// working on actions from script
	actions := strings.Split(script, "")
	actionsCount := len(actions)
	fmt.Printf("SR3# we have %d actions to do\n", actionsCount)

	toBeSentAction := Action3{}
	toBeSentAction.robotName = robotName

	if actionsCount == 0 {
		// nothing to do..
		// log <- "No action"
	} else {
		for _, act := range actions {
			fmt.Printf("SR3# executing action: %s\n", act)
			//
			switch act {
			case "A":
				toBeSentAction.action = 'A'
				break
			case "R":
				toBeSentAction.action = 'R'
				break
			case "L":
				toBeSentAction.action = 'L'
				break
			default:
				log <- "Unknown action"
			}
			// we must trigger the Room in order to update robot state, even if no
			// actions
			action <- toBeSentAction
		}
	}
	fmt.Println("SR3# closing action channel")
}