package robot

// const
const (
	N = Dir(0)
	S = Dir(1)
	W = Dir(2)
	E = Dir(3)
)

func (d Dir) String() string {
	switch d {
	case N:
		return "N"
	case S:
		return "S"
	case W:
		return "W"
	case E:
		return "E"
	}
	return "?"
}
