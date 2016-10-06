package checkers

import (
  "testing"
  "github.com/cl0ne/go-checkers/point"
)

func TestKill(t *testing.T) {
  checker := NewChecker(2, 2, true)
  checker.Kill()
  if checker.IsAlive() {
    t.Error("checker is alive after kill.")
  }
}

func TestNewChecker(t *testing.T) {
  checker_queen := NewChecker(3, 2, false)
  checker_alive := NewChecker(2, 4, true)
  checker_pos := NewChecker(4, 2, true)
  point := point.Point{ X:4, Y:2 }
  if checker_queen.IsQueen() {
    t.Error("checker is queen.")
  }
  if !checker_alive.IsAlive() {
    t.Error("checker was killed after creation.")
  }
  if checker_pos.Position() != point {
    t.Error("checker is not right position.")
  }
}

func TestMakeQueen(t *testing.T) {
  checker := NewChecker(3, 3, true)
  checker.MakeQueen()
  if !checker.IsQueen() {
    t.Error("checker is not queen.")
  }
}

func TestIsWhite(t *testing.T) {
  checker := NewChecker(3, 3, true)
  if checker.IsBlack() {
    t.Error("checker is black.")
  }
}

func TestIsBlack(t *testing.T) {
  checker := NewChecker(3, 3, false)
  if checker.IsWhite() {
    t.Error("checker is white.")
  }
}

func TestSetPosition(t *testing.T) {
  checker := NewChecker(4, 5, true)
  point := point.Point{ X:6, Y:2 }
  checker.SetPosition(point)
  if checker.Position() != point {
    t.Error("checker is not right position.")
  }
}
