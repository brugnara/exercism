package robot

import (
	"errors"
	"fmt"
)

// StartRobot3 starts the script for the given robot name
func StartRobot3(name, script string, action chan Action3, log chan string) {
	if name == "" {
		log <- "Invalid name"
	}
	for _, c := range script {
		// break on first invalid command
		if !isValidCommand(c) {
			log <- fmt.Sprint("Invalid command:", c)
			break
		}
		action <- Action3{Action(c), name}
	}

	// Tell the room this robot will not move again.
	action <- Action3{Action(0), name}
}

func isValidCommand(cmd rune) bool {
	return cmd == 'A' || cmd == 'L' || cmd == 'R'
}

func toHash(robots []Step3Robot, room *Rect) (hash map[string]int, err error) {
	hash = map[string]int{}
	occupied := map[Pos]bool{}
	for i, r := range robots {
		// out of room?
		if r.isOutOfRoom(room) {
			err = errors.New("Out of the room :/")
			return
		}

		// duplicated name
		if _, ok := hash[r.Name]; ok {
			// Bad
			err = errors.New("Duplicated name")
			return
		}
		hash[r.Name] = i

		// slot occupied?
		if _, ok := occupied[r.Pos]; ok {
			err = errors.New("slot occupied")
			return
		}
		occupied[r.Pos] = true
	}
	return
}

// Room3 is the room handler for step3
func Room3(
	extent Rect,
	robots []Step3Robot,
	action chan Action3,
	report chan []Step3Robot,
	log chan string) {
	// storing robots on a map
	hash, err := toHash(robots, &extent)
	if err != nil {
		log <- fmt.Sprint(err)
		report <- nil
		return
	}

	hasEnded := map[string]bool{}

	//
	go func() {
		for a := range action {
			// fmt.Println("Received action:", a)
			isAnEnd := a.Action == Action(0)
			if _, ok := hash[a.name]; !isAnEnd && !ok {
				log <- "Unknown robot:" + a.name
				continue
			}

			switch a.Action {
			// default as a fail-safe. Should never happens
			default:
				log <- "command unknown"
			case Action('A'):
				// fmt.Println("A")
				if !robots[hash[a.name]].moveForward(&extent, robots) {
					log <- "collision"
				}
			case Action('L'):
				// fmt.Println("L")
				robots[hash[a.name]].turnLeft()
			case Action('R'):
				// fmt.Println("R")
				robots[hash[a.name]].turnRight()
			case Action(0):
				// fmt.Printf("robot '%s' finished.\n", a.name)
				hasEnded[a.name] = true
				if len(hasEnded) == len(robots) {
					// fmt.Println("Closing Action")
					close(action)
				}
			}
		}
		report <- robots
	}()
}

func (robot Step3Robot) isOutOfRoom(room *Rect) bool {
	left := robot.Pos.Easting < room.Min.Easting
	right := robot.Pos.Easting > room.Max.Easting
	top := robot.Pos.Northing > room.Max.Northing
	bottom := robot.Pos.Northing < room.Min.Northing
	//
	return left || right || top || bottom
}

func (robot Step3Robot) samePlaced(robots []Step3Robot) bool {
	for _, r := range robots {
		if r.Name != robot.Name && r.Pos == robot.Pos {
			return true
		}
	}
	return false
}

// moveForward override
func (robot *Step3Robot) moveForward(room *Rect, robots []Step3Robot) bool {
	// cache
	x := robot.Pos.Easting
	y := robot.Pos.Northing

	// move
	switch robot.Dir {
	case N:
		robot.Pos.Northing++
	case S:
		robot.Pos.Northing--
	case E:
		robot.Pos.Easting++
	case W:
		robot.Pos.Easting--
	}

	// up-down
	robot.Pos.Northing = minRU(robot.Pos.Northing, room.Max.Northing)
	robot.Pos.Northing = maxRU(robot.Pos.Northing, room.Min.Northing)
	// <->
	robot.Pos.Easting = minRU(robot.Pos.Easting, room.Max.Easting)
	robot.Pos.Easting = maxRU(robot.Pos.Easting, room.Min.Easting)

	// check other robot positions!
	if robot.samePlaced(robots) {
		// rollback coords
		robot.Pos.Easting = x
		robot.Pos.Northing = y
	}

	// pos is different? Then we moved..
	return x != robot.Pos.Easting || y != robot.Pos.Northing
}
