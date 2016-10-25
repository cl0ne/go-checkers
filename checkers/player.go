package checkers

type Player struct {
	checkers           []Checker
	aliveCheckersCount int
	isWhite            bool
	availableMoves     map[*Checker][]Move
}

func newPlayer(checkersCount int, isWhite bool) *Player {
	checkers := make([]Checker, checkersCount)
	for i := range checkers {
		checkers[i].color = isWhite
		checkers[i].alive = true
	}
	availableMoves := make(map[*Checker][]Move)
	return &Player{
		checkers:           checkers,
		aliveCheckersCount: checkersCount,
		isWhite:            isWhite,
		availableMoves:     availableMoves,
	}
}

func (p Player) GetAliveCheckers() (alive []*Checker) {
	if p.aliveCheckersCount == 0 {
		return
	}
	for i := range p.checkers {
		c := &p.checkers[i]
		if c.IsAlive() {
			alive = append(alive, c)
		}
	}
	return
}

func (p Player) GetAvailabeCheckers() (available []*Checker) {
	for c := range p.availableMoves {
		available = append(available, c)
	}
	return
}

func (p Player) HasAliveCheckers() bool {
	return p.aliveCheckersCount != 0
}

func (p Player) HasAvailableMoves() bool {
	return len(p.availableMoves) != 0
}
