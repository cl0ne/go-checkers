package checkers

// Move describes move for checker
type Move struct {
	Target          Point
	CapturedChecker *Checker
	BecomeQueen     bool
}

// IsCapture reports whether the move caused capture.
func (m Move) IsCapture() bool {
	return m.CapturedChecker != nil
}

// CapturedPos returns position of captured checker for the move
// if it caused capture or point outside of the board otherwise.
func (m Move) CapturedPos() Point {
	if !m.IsCapture() {
		return Point{-1, -1}
	}
	return m.CapturedChecker.Position()
}
