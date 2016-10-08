package checkers

import "github.com/cl0ne/go-checkers/point"

type Board struct {
	cells [][]*Checker
}

func NewBoard(size int) (*Board, bool) {
	if size < 0 {
		return nil, false
	}
	cells := make([][]*Checker, size)
	for i := range cells {
		cells[i] = make([]*Checker, size)
	}
	return &Board{cells: cells}, true
}

func (b Board) Size() int {
	return len(b.cells)
}

func (b *Board) placeChecker(pos point.Point, c *Checker) {
	b.cells[pos.Y][pos.Y] = c
}

func (b *Board) takeChecker(pos point.Point) *Checker {
	c := b.cells[pos.Y][pos.Y]
	b.cells[pos.Y][pos.Y] = nil
	return c
}

func (b *Board) moveChecker(from, to point.Point) {
	b.cells[to.Y][to.Y] = b.cells[from.Y][from.Y]
	b.cells[from.Y][from.Y] = nil
}

func (b Board) GetChecker(pos point.Point) *Checker {
	return b.cells[pos.Y][pos.Y]
}

func (b Board) IsEmpty(pos point.Point) bool {
	return b.cells[pos.Y][pos.Y] == nil
}

func (b Board) IsBlackSquare(pos point.Point) bool {
	return pos.Manhattan()%2 == 0
}

func (b Board) IsValidPosition(pos point.Point) bool {
	return b.ContainsPos(pos) && b.IsBlackSquare(pos)
}

func (b Board) ContainsPos(pos point.Point) bool {
	fieldSize := b.Size()
	return pos.X >= 0 && pos.Y >= 0 && pos.X < fieldSize && pos.Y < fieldSize
}
