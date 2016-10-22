package checkers

import (
	"testing"
)

func TestKill(t *testing.T) {
	checker := newChecker(true)
	checker.kill()
	if checker.IsAlive() {
		t.Error("checker is alive after kill.")
	}
}

func TestNewChecker(t *testing.T) {
	checker_queen := newChecker(false)
	checker_alive := newChecker(true)
	checker_pos := newChecker(true)
	point := Point{X: 4, Y: 2}
	if checker_queen.IsQueen() {
		t.Error("checker is queen.")
	}
	if !checker_alive.IsAlive() {
		t.Error("checker was killed after creation.")
	}
	if checker_pos.Position() == point {
		t.Error("checker's coords must be zero, got ", point.String())
	}
}

func TestMakeQueen(t *testing.T) {
	checker := newChecker(true)
	checker.makeQueen()
	if !checker.IsQueen() {
		t.Error("checker is not queen.")
	}
}

func TestIsWhite(t *testing.T) {
	checker := newChecker(true)
	if checker.IsBlack() {
		t.Error("checker is black.")
	}
}

func TestIsBlack(t *testing.T) {
	checker := newChecker(false)
	if checker.IsWhite() {
		t.Error("checker is white.")
	}
}

func TestSetPosition(t *testing.T) {
	checker := newChecker(true)
	point := Point{X: 6, Y: 2}
	checker.setPosition(6, 2)
	if checker.Position() != point {
		t.Error("checker is not right position.")
	}
}
