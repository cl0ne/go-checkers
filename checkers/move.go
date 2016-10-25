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
