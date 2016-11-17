package checkers

// PlayerPoller is an interface for Game to get player selection of
// next move. Don't modify any of passed in arguments
type PlayerPoller interface {
	SelectChecker(availableCheckers []*Checker) int
	SelectTargetPos(availableMoves []Move) int
}
