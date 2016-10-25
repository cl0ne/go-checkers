package checkers

type Game struct {
	players      [2]*Player
	board        *Board
	activePlayer int
	isFinished   bool
}

func NewGame() *Game {
	boardSize := 8
	checkerCount := 12
	players := [...]*Player{
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

func (g Game) getOpponent() *Player {
	if g.activePlayer == 0 {
		return g.players[1]
	}
	return g.players[0]
}

func (g Game) getCurrentPlayer() *Player {
	return g.players[g.activePlayer]
}

var moveOffsets = [...]Point{
	Point{-1, 1}, Point{1, 1}, // forward
	Point{-1, -1}, Point{1, -1}, // back
}

// getAvailableMoves returns list of available moves for checker c.
// Checker must be alive and present on the game's board!
// Use capturesOnly to specify looking only for captures available to
// the checker or for any moves allowed by the game rules
func (g Game) getAvailableMoves(c *Checker, capturesOnly bool) (moves []Move) {
	if c == nil || !c.IsAlive() {
		return
	}

	pos := c.Position()
	board := g.board
	if board.GetChecker(pos.X, pos.Y) != c {
		return
	}

	lastRow := 0
	if c.IsWhite() {
		lastRow = board.LastRowIndex()
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

			if neighbour != nil {
				if captureFound != nil { // blocked
					break
				}

				if neighbour.IsWhite() == c.IsWhite() {
					break
				}

				captureFound = neighbour
				localPos = target

				continue
			}

			if captureFound == nil {
				if capturesOnly { // ignore moves
					if c.IsQueen() {
						continue // look for captures further
					}
					break
				}

				if !c.IsForwardMove(target) && !c.IsQueen() {
					break
				}
			} else if !capturesOnly {
				capturesOnly = true
				moves = make([]Move, 1)
			}

			becomeQueen := !c.IsQueen() && target.Y == lastRow
			moves = append(moves, Move{target, captureFound, becomeQueen})
			if !c.IsQueen() {
				break
			}
		}
	}
	return
}

func (g *Game) Start() {
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
	g.updatePlayerMoves(g.getCurrentPlayer())
}

func (g *Game) updatePlayerMoves(p *Player) {
	for _, c := range p.GetAliveCheckers() {
		moves := g.getAvailableMoves(c, false)
		if len(moves) > 0 {
			p.availableMoves[c] = moves
		}
	}
}
