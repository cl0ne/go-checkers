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
	checkers := []struct {
		checker Checker
		coordX, coordY int
	}{
		{ newChecker(false), 4, 5 },
		{ newChecker(true), 5, 6 },
		{ newChecker(false), 3, 6 },
		{ newChecker(true), 3, 4 },
		{ newChecker(false), 5, 4 },
		{ newChecker(true), 6, 7 },
		{ newChecker(false), 2, 7 },
		{ newChecker(true), 2, 3 },
		{ newChecker(false), 6, 3 },
	}
	cases_range := []struct {
		checker Checker
		coordX, coordY int
		moves []Move
	}{
		{ newChecker(true), -3, -5, nil },
		{ newChecker(false), 10, 10, nil },
	}
	cases_threesome := []struct {
		one, two, three Checker
		coordOne, coordTwo, coordThree Point
	}{
		{ newChecker(true), newChecker(false), newChecker(false),
		Point{ X:4, Y:4 }, Point{ X:5, Y:5 }, Point{ X:6, 6} },
	}
	cases_backmove := []struct {
		one, two Checker
		coordOne, coordTwo Point
		moves []Move
	}{
		{ newChecker(true, newChecker(false), Point{ X:4, Y:5 }, Point{ X:5, Y:4 },
	 	nil },

	}

    game := NewGame()
    board := game.GetBoard()

	for _, c := range cases_range {
		board.placeChecker(c.coordX, c.coordY, &c.checker)
		moves := game.getAvailableMoves(&c.checker, false)
		if moves != c.moves {
			t.Error("Checker at ", c.checker.Position(), " shouldn't have moves.")
		}
	}

	for _, c := range cases_threesome {
		board.placeChecker(c.coordOne.X, c.coordOne.Y, &c.one)
		board.placeChecker(c.coordTwo.X, c.coordTwo.Y, &c.two)
		board.placeChecker(c.coordThree.X, c.coordThree.Y,&c.three)
		moves := game.getAvailableMoves(&c.one, true)
		if moves != nil {
			t.Error("Nothing to kill from ", c.coordOne)
		}
		moves = game.getAvailableMoves(&c.one, false)
		if moves == nil {
			t.Error("Checker at ", c.coordOne, " can move without capture.")
		}
		board.takeChecker(c.coordOne.X, c.coordOne.Y)
		board.takeChecker(c.coordTwo.X, c.coordTwo.Y)
		board.takeChecker(c.coordThree.X, c.coordThree.Y)
	}

	for _, c := range cases_backmove {
		board.placeChecker(c.coordOne.X, c.coordOne.Y, &c.one)
		board.placeChecker(c.coordTwo.X, c.coordTwo.Y, &c.two)
		moves := game.getAvailableMoves(&c.one, true)
		if moves != c.moves {
			t.Error("Checker shouldn't move or kill back.")
		}
		board.takeChecker(c.coordOne.X, c.coordOne.Y)
		board.takeChecker(c.coordTwo.X, c.coordTwo.Y)
	}

	for _, c := checkers {
		board.placeChecker(c.coordX, c.coordY, &c.checker)
	}

}
