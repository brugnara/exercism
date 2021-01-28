package robot

// moving Step1Robot

/*
Advance moves the robot to the direction
  N
W + E
  S
*/
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

/*
Right turns to right
  N
W + E
  S
*/
func Right() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = E
	case E:
		Step1Robot.Dir = S
	case S:
		Step1Robot.Dir = W
	case W:
		Step1Robot.Dir = N
	}
}

/*
Left turns to left
  N
W + E
  S
*/
func Left() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = W
	case W:
		Step1Robot.Dir = S
	case S:
		Step1Robot.Dir = E
	case E:
		Step1Robot.Dir = N
	}
}
