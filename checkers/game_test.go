package checkers

import (
	"sort"
	"testing"
)

type ByTarget []Move

func (m ByTarget) Len() int {
	return len(m)
}

func (m ByTarget) Less(i, j int) bool {
	return m[i].Target.Less(m[j].Target)
}

func (m ByTarget) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func CompareMoves(moves, expected []Move) bool {
	sort.Sort(ByTarget(moves))
	sort.Sort(ByTarget(expected))
	for i := range moves {
		if moves[i] != expected[i] {
			return false
		}
	}
	return true
}

func TestNewGame(t *testing.T) {
	game := NewGame()
	board := game.GetBoard()
	if board == nil {
		t.Error("Board wasn't created.")
	}
}

func TestGetAvailableMoves(t *testing.T) {
	checkers := []*struct {
		ch   Checker
		X, Y int
	}{
		{newChecker(false), 4, 4},
		{newChecker(true), 5, 5},
		{newChecker(false), 3, 5},
		{newChecker(true), 3, 3},
		{newChecker(false), 5, 3},
		{newChecker(true), 6, 6},
		{newChecker(false), 2, 6},
		{newChecker(true), 2, 2},
		{newChecker(false), 6, 2},
	}

	expectedMoves := [][]Move{
		nil,
		{Move{Target: Point{X: 4, Y: 6}}},
		{Move{Target: Point{X: 2, Y: 4}}},
		{Move{Target: Point{X: 2, Y: 4}}},
		{Move{Target: Point{X: 4, Y: 2}}},
		{
			Move{Target: Point{X: 7, Y: 7}, BecomeQueen: true},
			Move{Target: Point{X: 5, Y: 7}, BecomeQueen: true},
		},
		{Move{Target: Point{X: 1, Y: 5}}},
		{Move{Target: Point{X: 1, Y: 3}}},
		{
			Move{Target: Point{X: 5, Y: 1}},
			Move{Target: Point{X: 7, Y: 1}},
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
		if len(moves) == len(expectedMoves[i]) {
			if !CompareMoves(moves, expectedMoves[i]) {
				t.Error("Moves got: ", moves, " expected: ", expectedMoves[i])
			}
		}
	}

	if t.Failed() {
		t.Logf("Current board state:\n%s", board.DebugString())
	}
}
