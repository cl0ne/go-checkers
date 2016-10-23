package checkers

import (
	"testing"
)

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

func TestCheckerOperations(t *testing.T) {
	blackChecker := newChecker(false)
	whiteChecker := newChecker(true)
	t.Run("MakeQueen", func(t *testing.T) {
		whiteChecker.makeQueen()
		if !whiteChecker.IsQueen() {
			t.Error("checker is not queen.")
		}
	})

	t.Run("IsWhite", func(t *testing.T) {
		if whiteChecker.IsBlack() {
			t.Error("checker is black.")
		}
	})

	t.Run("IsBlack", func(t *testing.T) {
		if blackChecker.IsWhite() {
			t.Error("checker is white.")
		}
	})

	t.Run("SetPosition", func(t *testing.T) {
		point := Point{X: 6, Y: 2}
		whiteChecker.setPosition(6, 2)
		if whiteChecker.Position() != point {
			t.Error("checker is not right position.")
		}
	})

	t.Run("Kill", func(t *testing.T) {
		blackChecker.kill()
		if blackChecker.IsAlive() {
			t.Error("checker is alive after kill.")
		}
	})
}
