package checkers

import (
	"fmt"
	"math"
	"strings"
)

/*
Checker board class.

Example board:

   abcdefgh
  ..........
 8| # # # #|8
 7|# # # # |7
 6| # # # #|6
 5|# # # # |5
 4| # # # #|4
 3|# # # # |3
 2| # # # #|2
 1|# # # # |1
  ''''''''''
   abcdefgh

Here a1 has position (0,0) and h8 is on (7,7)
*/
type Board struct {
	cells [][]*Checker
}

func NewBoard(size int) *Board {
	if size < 0 {
		return nil
	}
	cells := make([][]*Checker, size)
	for i := range cells {
		cells[i] = make([]*Checker, size)
	}
	return &Board{cells: cells}
}

func (b Board) Size() int {
	return len(b.cells)
}

var checkerDebug = map[bool]map[bool]string{
	false: {false: "b", true: "B"},
	true:  {false: "w", true: "W"},
}

func (b Board) DebugString() string {
	size := b.Size()
	rowNumberFormat := fmt.Sprintf("%%-%dd|", int(math.Floor(math.Log10(2)+1)))
	rows := make([]string, size+1)
	for r := 0; r < size; r++ {
		cols := make([]string, size+2)
		cols[0] = fmt.Sprintf(rowNumberFormat, r)
		cols[size+1] = "|"
		for c := 0; c < size; c++ {
			cell := " "
			ch := b.GetChecker(c, r)
			if ch != nil {
				cell = checkerDebug[ch.IsWhite()][ch.IsQueen()]
			} else if !b.IsWhiteSquare(Point{c, r}) {
				cell = "."
			}
			cols[c+1] = cell
		}
		rows[size-r-1] = strings.Join(cols, "")
	}
	return strings.Join(rows, "\n")
}

func (b *Board) clear() {
	for _, r := range b.cells {
		for c := range r {
			r[c] = nil
		}
	}
}

func (b *Board) placeChecker(x, y int, c *Checker) bool {
	if !b.ContainsPos(x, y) {
		return false
	}

	b.cells[y][x] = c
	if c != nil {
		c.setPosition(x, y)
	}
	return true
}

func (b *Board) takeChecker(x, y int) *Checker {
	c := b.cells[y][x]
	b.cells[y][x] = nil
	return c
}

func (b *Board) moveChecker(from, to Point) bool {
	if from == to {
		return false
	}
	if !b.ContainsPos(from.X, from.Y) || !b.ContainsPos(to.X, to.Y) {
		return false
	}
	if !b.IsEmpty(to.X, to.Y) {
		return false
	}
	c := b.takeChecker(from.X, from.Y)
	return b.placeChecker(to.X, to.Y, c)
}

func (b Board) GetChecker(x, y int) *Checker {
	return b.cells[y][x]
}

func (b Board) IsEmpty(x, y int) bool {
	if !b.ContainsPos(x, y) {
		return true
	}
	return b.cells[y][x] == nil
}

func (b Board) IsBlackSquare(pos Point) bool {
	return pos.Manhattan()%2 == 0
}

func (b Board) IsWhiteSquare(pos Point) bool {
	return pos.Manhattan()%2 == 1
}

func (b Board) ContainsPos(x, y int) bool {
	fieldSize := b.Size()
	return x >= 0 && y >= 0 && x < fieldSize && y < fieldSize
}

func (b Board) LastRowIndex() int {
	return b.Size() - 1
}

func (b Board) LastColumnIndex() int {
	return b.Size() - 1
}
