package checkers

/*
Checker board class
For example:

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

func (b *Board) placeChecker(x, y int, c *Checker) {
	b.cells[y][x] = c
	if c != nil {
		c.setPosition(x, y)
	}
}

func (b *Board) takeChecker(x, y int) *Checker {
	c := b.cells[y][x]
	b.cells[y][x] = nil
	return c
}

func (b *Board) moveChecker(from, to Point) {
	if from == to {
		return
	}
	if !b.ContainsPos(from.X, from.Y) || !b.ContainsPos(to.X, to.Y) {
		return
	}
	c := b.takeChecker(from.X, from.Y)
	b.placeChecker(to.X, to.Y, c)
}

func (b Board) GetChecker(x, y int) *Checker {
	return b.cells[y][x]
}

func (b Board) IsEmpty(x, y int) bool {
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

