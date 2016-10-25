package checkers

import "testing"

func TestNewGame(t *testing.T) {
	game := NewGame()
	board := game.GetBoard()
	if board == nil {
		t.Error("Board wasn't created.")
	}
}

func TestGetAvailableMoves(t *testing.T) {
	checkers := []Checker{newChecker(true),
		newChecker(false), newChecker(true),
	}
	game := NewGame()
	board := game.GetBoard()
	board.placeChecker(4, 5, &checkers[0])
	moves := game.getAvailableMoves(&checkers[0], false)
	if moves == nil {
		t.Error("Checker at ", checkers[0].Position(), " has no moves.")
	}
	board.placeChecker(-4, -2, &checkers[1])
	moves = game.getAvailableMoves(&checkers[1], false)
	if moves != nil {
		t.Error("Checker at ", checkers[1].Position(), " has moves instead zero.")
	}
	board.placeChecker(10, 10, &checkers[2])
	moves = game.getAvailableMoves(&checkers[2], false)
	if moves != nil {
		t.Error("Checker at ", checkers[2].Position(), " has moves instead zero.")
	}
}
