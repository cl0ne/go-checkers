package checkers

import "fmt"

type Checker struct {
	position Point
	color    bool
	queen    bool
	alive    bool
}

func newChecker(isWhite bool) Checker {
	return Checker{
		color: isWhite,
		queen: false,
		alive: true,
	}
}

func (ch Checker) Position() Point {
	return ch.position
}

func (ch *Checker) setPosition(x, y int) {
	ch.position.X = x
	ch.position.Y = y
}

func (c Checker) IsForwardMove(target Point) bool {
	return c.Position() != target && c.IsWhite() == (c.Position().Y < target.Y)
}

func (ch Checker) IsWhite() bool {
	return ch.color
}

func (ch Checker) IsBlack() bool {
	return !ch.color
}

func (ch Checker) IsQueen() bool {
	return ch.queen
}

func (ch *Checker) makeQueen() {
	ch.queen = true
}

func (ch Checker) IsAlive() bool {
	return ch.alive
}

var stateMap = map[string]map[bool]string{
	"alive": {true: "alive", false: "dead"},
	"color": {true: "white", false: "black"},
	"queen": {true: " queen", false: ""},
}

func (ch Checker) String() string {
	return fmt.Sprintf("<%v %v%v %v>",
		stateMap["alive"][ch.IsAlive()],
		stateMap["color"][ch.IsWhite()],
		stateMap["queen"][ch.IsQueen()],
		ch.Position(),
	)
}

func (ch *Checker) kill() {
	ch.alive = false
}
