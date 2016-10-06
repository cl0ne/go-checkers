package checkers

import "github.com/cl0ne/go-checkers/point"

type Checker struct {
	position point.Point
	color    bool
	queen    bool
	alive    bool
}

func NewChecker(x, y int, isWhite bool) Checker {
	return Checker{
		position: point.Point{X: x, Y: y},
		color:    isWhite, queen: false, alive: true}
}

func (ch Checker) Position() point.Point {
	return ch.position
}

func (ch *Checker) SetPosition(p point.Point) {
	ch.position = p
}

func (ch Checker) IsWhite() bool {
	return ch.color
}

func (ch Checker) IsBlack() bool {
	return !ch.color
}

func (ch Checker) IsQueen() bool {
	return ch.queen
}

func (ch *Checker) MakeQueen() {
	ch.queen = true
}

func (ch Checker) IsAlive() bool {
	return ch.alive
}

func (ch *Checker) Kill() {
	ch.alive = false
}
