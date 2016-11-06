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
	// TODO: Fix positions (checkers are placed on white cells!)
	checkers := []*struct {
		ch   Checker
		X, Y int
	}{
		{newChecker(false), 4, 5},
		{newChecker(true), 5, 6},
		{newChecker(false), 3, 6},
		{newChecker(true), 3, 4},
		{newChecker(false), 5, 4},
		{newChecker(true), 6, 7},
		{newChecker(false), 2, 7},
		{newChecker(true), 2, 3},
		{newChecker(false), 6, 3},
	}

	expectedMoves := [][]Move{
		nil,
		{Move{Target: Point{X: 4, Y: 7}, BecomeQueen: true}},
		{Move{Target: Point{X: 2, Y: 5}}},
		{Move{Target: Point{X: 2, Y: 5}}},
		{Move{Target: Point{X: 4, Y: 3}}},
		nil,
		{Move{Target: Point{X: 1, Y: 6}}},
		{Move{Target: Point{X: 1, Y: 4}}},
		{
			Move{Target: Point{X: 5, Y: 2}},
			Move{Target: Point{X: 7, Y: 2}},
		},
	}

	game := NewGame()
	board := game.GetBoard()

	for _, c := range checkers {
		moves := game.getAvailableMoves(&c.ch, false)
		if len(moves) > 0 {
			t.Error("Checker at", c.ch.Position(), "is not on board but has moves:", moves)
		}
	}

	for _, c := range checkers {
		board.placeChecker(c.X, c.Y, &c.ch)
	}

	for i, c := range checkers {
		moves := game.getAvailableMoves(&c.ch, false)
		// we should compare sorted moves (both expected and acquired from game)
		// since order can be different. First we have to compare lengths of
		// both slices and only the do in-depth compare (i.e. sort and compare
		// elements). It's requied to extract comparison code ino separate
		// function.
		if !reflect.DeepEqual(moves, expectedMoves[i]) {
			if expectedMoves[i] == nil {
				t.Error("Checker at", c.ch.Position(), "shouldn't have moves.")
			} else {
				t.Error("Checker at", c.ch.Position(), "should have at least 1 move.")
			}
		}
	}

	if t.Failed() {
		t.Logf("Current board state:\n%s", board.DebugString())
	}
}
