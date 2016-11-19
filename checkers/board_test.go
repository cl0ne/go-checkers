package checkers

import "testing"

func TestNewBoard(t *testing.T) {
	uselessBoard := NewBoard(0)
	if uselessBoard == nil {
		t.Error("Zero-sized boards aren't very useful, but got nil.")
	}
	if uselessBoard.Size() != 0 {
		t.Error("Zero-sized board after creation got size of", uselessBoard.Size())
	}

	invalidBoard := NewBoard(-1)
	if invalidBoard != nil {
		t.Error("Boards of negative size not allowed, got size of", invalidBoard.Size())
	}

	board8x8 := NewBoard(8)
	if board8x8 == nil {
		t.Error("Can't create 8x8 board got nil.")
	}
	if board8x8.Size() != 8 {
		t.Error("Size of created 8x8 board is", board8x8.Size())
	}
}

func TestBoardOperations(t *testing.T) {
	board := NewBoard(8)
	t.Run("IsEmpty", func(t *testing.T) {
		checker := newChecker(true)
		p := Point{X: 3, Y: 3}
		board.placeChecker(p.X, p.Y, &checker)
		if board.IsEmpty(p.X, p.Y) {
			t.Error("Cell", p, "shouldn't be empty.")
		}
		p.X, p.Y = 4, 6
		if !board.IsEmpty(p.X, p.Y) {
			t.Error("Cell", p, "should be empty.")
		}
	})

	t.Run("PlaceChecker", func(t *testing.T) {
		cases := []*struct {
			c      Checker
			pos    Point
			result bool
		}{
			{newChecker(true), Point{X: -2, Y: -5}, false},
			{newChecker(false), Point{X: 10, Y: 10}, false},
			{newChecker(true), Point{X: 5, Y: 6}, true},
		}
		for _, c := range cases {
			ok := board.placeChecker(c.pos.X, c.pos.Y, &c.c)
			if ok != c.result {
				t.Error("Checker placement at", c.pos, "gave unexpected:", ok)
				continue
			}
			if !ok && !board.IsEmpty(c.pos.X, c.pos.Y) {
				t.Error("Place at ", c.pos, "should be empty: checker is not placed")
			}
			if ok {
				if c.c.Position() != c.pos {
					t.Error("Checker position should be set to ", c.pos)
				}
				if &c.c != board.GetChecker(c.pos.X, c.pos.Y) {
					t.Error("Place at ", c.pos, "should contain placed checker")
				}
			}
		}
	})

	t.Run("ContainsPos", func(t *testing.T) {
		boardSize := board.Size()
		for x := 0; x < board.Size(); x++ {
			for y := 0; y < board.Size(); y++ {
				contains := board.ContainsPos(x, y)
				if !contains {
					t.Error("Cell", Point{X: x, Y: y}, "is within the board.")
				}
			}
		}
		outside := []Point{
			Point{-1, -1},
			Point{0, -1},
			Point{-1, 0},
			Point{boardSize, boardSize},
			Point{boardSize, 0},
			Point{0, boardSize},
		}
		for _, p := range outside {
			contains := board.ContainsPos(p.X, p.Y)
			if contains {
				t.Error("Cell", p, "is outside of the board.")
			}
		}
	})

	whiteCell := Point{X: 2, Y: 1}
	blackCell := Point{X: 4, Y: 0}

	t.Run("IsBlackSquare", func(t *testing.T) {
		if !board.IsBlackSquare(blackCell) {
			t.Error("Cell", blackCell, "must be black.")
		}
		if board.IsBlackSquare(whiteCell) {
			t.Error("Cell", whiteCell, "must be white.")
		}
	})

	t.Run("IsWhiteSquare", func(t *testing.T) {
		if !board.IsWhiteSquare(whiteCell) {
			t.Error("Cell", whiteCell, "must be white.")
		}
		if board.IsWhiteSquare(blackCell) {
			t.Error("Cell", blackCell, "must be black.")
		}
	})

	t.Run("TakeChecker", func(t *testing.T) {
		p := Point{X: 2, Y: 2}
		checker := newChecker(true)
		board.placeChecker(p.X, p.Y, &checker)
		c := board.takeChecker(p.X, p.Y)
		if c == nil {
			t.Error("Can't take checker at", p, ": cell is empty.")
		}
	})

	t.Run("GetChecker", func(t *testing.T) {
		p := Point{X: 2, Y: 4}
		checker := newChecker(false)
		board.placeChecker(p.X, p.Y, &checker)
		c := board.GetChecker(p.X, p.Y)
		if c == nil {
			t.Error("Checker at", p, "is missing.")
		}
		p.X, p.Y = 5, 5
		c = board.GetChecker(p.X, p.Y)
		if c != nil {
			t.Error("Cell", p, "should be empty")
		}
	})

	t.Run("MoveChecker", func(t *testing.T) {
		from := Point{X: 3, Y: 4}
		checker := newChecker(true)
		checker2 := newChecker(true)
		to := Point{X: 4, Y: 3}
		board.placeChecker(from.X, from.Y, &checker)
		ok := board.moveChecker(from, to)
		if !ok {
			t.Error("Checker wasn't moved from", from, "to", to)
		}
		from = to
		to.X, to.Y = -3, -6
		ok = board.moveChecker(from, to)
		if ok {
			t.Error("Invalid position at ", to, ". Out of range of board.")
		}
		to.X, to.Y = 10, 8
		board.moveChecker(from, to)
		if ok {
			t.Error("Invalid position at ", to, ". Out of range of board.")
		}
		to = from
		ok = board.moveChecker(from, to)
		if ok {
			t.Error("Checker was moved on the same point ", to)
		}
		to.X, to.Y = 5, 4
		board.placeChecker(to.X, to.Y, &checker2)
		ok = board.moveChecker(from, to)
		if ok {
			t.Error("Checker at ", from, " shouldn't move to occupied ", to)
		}
	})
}
