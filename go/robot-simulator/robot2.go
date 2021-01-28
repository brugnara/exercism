package robot

func StartRobot(cmd chan Command, act chan Action) {
	go func() {
		defer close(act)

		for c := range cmd {
			act <- Action(c)
		}
	}()
}

func Room(extent Rect, robot Step2Robot, act chan Action, rep chan Step2Robot) {
	go func() {
		defer close(rep)
		for a := range act {
			operateInto(&extent, &robot, a)
		}
		rep <- robot
	}()
}

func operateInto(room *Rect, robot *Step2Robot, a Action) {
	switch a {
	case 'A':
		moveForward(room, robot)
	case 'L':
		turnLeft(robot)
	case 'R':
		turnRight(robot)
	}
}

func turnLeft(robot *Step2Robot) {
	switch robot.Dir {
	case N:
		robot.Dir = W
	case W:
		robot.Dir = S
	case S:
		robot.Dir = E
	case E:
		robot.Dir = N
	}
}

func turnRight(robot *Step2Robot) {
	switch robot.Dir {
	case N:
		robot.Dir = E
	case E:
		robot.Dir = S
	case S:
		robot.Dir = W
	case W:
		robot.Dir = N
	}
}

func moveForward(room *Rect, robot *Step2Robot) {
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
	// check
	// up-down
	robot.Pos.Northing = minRU(robot.Pos.Northing, room.Max.Northing)
	robot.Pos.Northing = maxRU(robot.Pos.Northing, room.Min.Northing)
	// <->
	robot.Pos.Easting = minRU(robot.Pos.Easting, room.Max.Easting)
	robot.Pos.Easting = maxRU(robot.Pos.Easting, room.Min.Easting)
}
