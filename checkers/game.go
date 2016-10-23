package checkers

type Game struct {
	players      [2]Player
	board        *Board
	activePlayer int
	isFinished   bool
}

func NewGame() *Game {
	boardSize := 8
	checkerCount := 12
	players := [...]Player{
		newPlayer(checkerCount, true),
		newPlayer(checkerCount, false),
	}
	board := NewBoard(boardSize)
	if board == nil {
		return nil
	}
	return &Game{players, board, 0, false}
}

func (g Game) GetBoard() *Board {
	return g.board
}

var moveOffsets = [...]Point{
	Point{-1, 1}, Point{1, 1}, // forward
	Point{-1, -1}, Point{1, -1}, // back
}

func (g Game) getAvailableMoves(c *Checker) (moves []Move) {
	if c == nil || !c.IsAlive() ||
		g.board.GetChecker(c.Position().X, c.Position().Y) != c {
		return
	}
	pos := c.Position()
	board := g.board
	hasCaptures := false
	lastRow := 0
	if c.IsWhite() {
		lastRow = g.board.LastRowIndex()
	}
	for _, offset := range moveOffsets {
		var captureFound *Checker = nil
		localPos := pos
		for {
			target := localPos.Add(offset)
			if !board.ContainsPos(target.X, target.Y) {
				break
			}

			neighbour := board.GetChecker(target.X, target.Y)

			if neighbour == nil {
				if captureFound == nil {
					if hasCaptures { // then ignore moves
						if !c.IsQueen() {
							break
						}
						continue // look only for captures
					}

					if !c.IsForwardMove(target) && !c.IsQueen() {
						break
					}
				} else if !hasCaptures {
					hasCaptures = true
					moves = make([]Move, 1)
				}

				becomeQueen := !c.IsQueen() && target.Y == lastRow
				moves = append(moves, Move{target, captureFound, becomeQueen})
				if !c.IsQueen() {
					break
				}
			}

			if captureFound != nil { // blocked
				break
			}

			if neighbour.IsWhite() == c.IsWhite() {
				break
			}

			captureFound = neighbour
			localPos = target
		}
	}
	return
}

func (g *Game) StartGame() {
	g.isFinished = false
	g.activePlayer = 0

	boardSize := g.board.Size()
	occupiedRows := boardSize/2 - 1
	lastRow, lastColumn := g.board.LastRowIndex(), g.board.LastColumnIndex()
	for r, i := 0, 0; r < occupiedRows; r++ {
		for c := r % 2; c < boardSize; c += 2 {
			checker := &g.players[0].checkers[i]
			g.board.placeChecker(c, r, checker)
			checker = &g.players[1].checkers[i]
			g.board.placeChecker(lastColumn-c, lastRow-r, checker)
			i++
		}
	}

	for _, p := range g.players {
		for i := range p.checkers {
			c := &p.checkers[i]
			moves := g.getAvailableMoves(c)
			if len(moves) > 0 {
				p.availableMoves[c] = moves
			}
		}
	}
}
