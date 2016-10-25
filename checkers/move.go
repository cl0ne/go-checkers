package checkers

type Move struct {
	Target          Point
	CapturedChecker *Checker
	BecomeQueen     bool
}

func (m Move) IsCapture() bool {
	return m.CapturedChecker != nil
}
