package checkers

type Checker struct {
	position Point
	color    bool
	queen    bool
	alive    bool
}

func newChecker(x, y int, isWhite bool) Checker {
	return Checker{
		position: Point{X: x, Y: y},
		color:    isWhite,
		queen:    false,
		alive:    true,
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

func (ch *Checker) kill() {
	ch.alive = false
}