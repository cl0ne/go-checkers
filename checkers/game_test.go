package checkers

import (
	"reflect"
	"testing"
)

func TestNewGame(t *testing.T) {
	game := NewGame()
	board := game.GetBoard()
	if board == nil {
		t.Error("Board wasn't created.")
	}
}

func TestGetAvailableMoves(t *testing.T) {
	checkers := []*struct {
		checker Checker
		X, Y    int
		moves   []Move
	}{
		{newChecker(false), 4, 5, nil},
		{newChecker(true), 5, 6, []Move{
			Move{Target: Point{X: 4, Y: 7}, BecomeQueen: true},
		}},
		{newChecker(false), 3, 6, []Move{
			Move{Target: Point{X: 2, Y: 5}},
		}},
		{newChecker(true), 3, 4, []Move{
			Move{Target: Point{X: 2, Y: 5}},
		}},
		{newChecker(false), 5, 4, []Move{
			Move{Target: Point{X: 4, Y: 3}},
		}},
		{newChecker(true), 6, 7, nil},
		{newChecker(false), 2, 7, []Move{
			Move{Target: Point{X: 1, Y: 6}},
		}},
		{newChecker(true), 2, 3, []Move{
			Move{Target: Point{X: 1, Y: 4}},
		}},
		{newChecker(false), 6, 3, []Move{
			Move{Target: Point{X: 5, Y: 2}},
			Move{Target: Point{X: 7, Y: 2}},
		}},
	}
	cases_backmove := []struct {
		checker Checker
		X, Y    int
		moves   []Move
	}{
		{newChecker(true), 4, 2, []Move{
			Move{Target: Point{X: 3, Y: 3}},
			Move{Target: Point{X: 5, Y: 3}},
		}},
		{newChecker(false), 1, 5, []Move{
			Move{Target: Point{X: 0, Y: 4}},
			Move{Target: Point{X: 2, Y: 4}},
		}},
	}
	cases_range := []struct {
		checker        Checker
		coordX, coordY int
	}{
		{newChecker(true), -3, -5},
		{newChecker(false), 10, 10},
	}

	game := NewGame()
	board := game.GetBoard()

	for _, c := range cases_range {
		board.placeChecker(c.coordX, c.coordY, &c.checker)
		moves := game.getAvailableMoves(&c.checker, false)
		if moves != nil {
			t.Error("Checker at ", c.checker.Position(), " shouldn't have moves.")
		}
	}

	for _, c := range cases_backmove {
		board.placeChecker(c.X, c.Y, &c.checker)
		moves := game.getAvailableMoves(&c.checker, false)
		if !reflect.DeepEqual(moves, c.moves) {
			t.Error("Checker can move only forward.")
		}
	}

	for _, c := range checkers {
		board.placeChecker(c.X, c.Y, &c.checker)
	}

	for _, i := range checkers {
		moves := game.getAvailableMoves(&i.checker, false)
		if !reflect.DeepEqual(moves, i.moves) {
			if i.moves == nil {
				t.Error("Checker at (", i.X, ",", i.Y, ") shouldn't have moves.")
			} else {
				t.Error("Checker at (", i.X, ",", i.Y, ") should have at least 1 move.")
			}
		}
	}

	if t.Failed() {
		t.Logf("Current board state:\n%s", board.DebugString())
	}
}
