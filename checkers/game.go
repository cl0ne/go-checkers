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

func (g Game) IsFinished() bool {
	return g.isFinished
}

func (g Game) IsWhitesMove() bool {
	return g.activePlayer == 0
}

func (g Game) IsBlacksMove() bool {
	return g.activePlayer == 1
}

func (g Game) IsBlackWin() bool {
	return g.IsFinished() && g.IsWhitesMove()
}

func (g Game) IsWhiteWin() bool {
	return g.IsFinished() && g.IsBlacksMove()
}

func (g Game) getOpponent() *Player {
	return g.players[g.getOpponentIndex()]
}

func (g Game) getOpponentIndex() int {
	return g.activePlayer ^ 1
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
		target := pos
		for {
			target = target.Add(offset)

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
				moves = make([]Move, 0, 1)
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
	p.clearAvailableMoves()

	capturesOnly := false
	for _, c := range p.GetAliveCheckers() {
		moves := g.getAvailableMoves(c, capturesOnly)

		if len(moves) > 0 {
			if !capturesOnly && moves[0].IsCapture() {
				p.clearAvailableMoves()
				capturesOnly = true
			}
			p.availableMoves[c] = moves
		}
	}
}

func dupMoves(moves []Point) (clone []Point) {
	clone = make([]Point, len(moves))
	copy(clone, moves)
	return
}

func (g *Game) Update() {
	if g.IsFinished() {
		return
	}
	player := g.getCurrentPlayer()
	opponent := g.getOpponent()

	availableCheckers := player.GetAvailabeCheckers()
	checkerIndex := 0 // UI.SelectChecker(availableCheckers)
	selectedChecker := availableCheckers[checkerIndex]

	availableMoves := player.availableMoves[selectedChecker]

	for {
		selectedMove := 0 // UI.SelectTargetPos(dupMoves(availableMoves))
		move := availableMoves[selectedMove]

		g.board.moveChecker(selectedChecker.Position(), move.Target)

		if move.BecomeQueen {
			selectedChecker.makeQueen()
		}

		if !move.IsCapture() {
			break
		}

		opponent.aliveCheckersCount--
		capturePos := move.CapturedPos()
		captured := g.board.takeChecker(capturePos.X, capturePos.Y)
		captured.kill()

		captures := g.getAvailableMoves(selectedChecker, true)
		if len(captures) == 0 {
			break
		}
		availableMoves = captures
	}

	player.clearAvailableMoves()
	g.updatePlayerMoves(opponent)
	g.isFinished = !(opponent.HasAliveCheckers() && opponent.HasAvailableMoves())
	g.activePlayer = g.getOpponentIndex()
}
