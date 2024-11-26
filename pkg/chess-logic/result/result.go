package result

type Result byte

const (
	Win Result = iota
	Lose
	Draw
	Unknown
)

func (r Result) String() string {
	switch r {
	case Win:
		return "win"
	case Lose:
		return "lose"
	case Draw:
		return "draw"
	case Unknown:
		return "unknown"
	}
	panic("unknown result")
}
