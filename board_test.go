package checkers

import (
  "testing"
  "github.com/cl0ne/go-checkers/point"
)

func TestNewBoard(t *testing.T) {
  board_zero := NewBoard(0)
  if board_zero != nil {
    t.Error("Size: ", board_zero.Size(), ". Board's size shouldn't be zero.")
  }
  board_neg := NewBoard(-1)
  if board_neg != nil {
    t.Error("Size: ", board_neg.Size(), ". Board's size shouldn't be negative.")
  }
}

func TestIsEmpty(t *testing.T) {
  board := NewBoard(8)
  checker := newChecker(3, 3, true)
  point_check := point.Point{ X:3, Y:3 }
  board.placeChecker(3, 3, &checker)
  if board.IsEmpty(3, 3) {
    t.Error("Cell at ", point_check.String(), " is empty.")
  }
  point_check.X = 4
  point_check.Y = 6
  if !board.IsEmpty(4, 6) {
    t.Error("Cell at ", point_check.String(), " is not empty.")
  }
}

func TestContainsPos(t *testing.T) {
  board := NewBoard(6)
  point := point.Point{ X:7, Y:8 }
  ok_pos := board.ContainsPos(7, 8)
  if ok_pos {
    t.Error("Invalid position at ", point.String(), ". Out of range of board.")
  }
  point.X, point.Y = 3, 2
  ok_pos = board.ContainsPos(3, 2)
  if !ok_pos {
    t.Error("Valid position at ", point.String(), ". In range of board.")
  }
}

func TestIsBlackSquare(t *testing.T) {
  board := NewBoard(8)
  point_black := point.Point{ X:4, Y:0 }
  point_white := point.Point{ X:2, Y:1 }
  if !board.IsBlackSquare(point_black) {
    t.Error("Cell at ", point_black.String(), " must be black.")
  }
  if board.IsBlackSquare(point_white) {
    t.Error("Cell at ", point_white.String(), " must be white.")
  }
}

func TestIsWhiteSquare(t *testing.T) {
  board := NewBoard(8)
  point_white := point.Point{ X:2, Y:1 }
  point_black := point.Point{ X:4, Y:0 }
  if !board.IsWhiteSquare(point_white) {
    t.Error("Cell at ", point_white.String(), " must be white.")
  }
  if board.IsWhiteSquare(point_black) {
    t.Error("Cell at ", point_black.String(), " must be black.")
  }
}

func TestTakeChecker(t *testing.T) {
  board := NewBoard(6)
  checker := newChecker(2, 2, true)
  point := point.Point{ X:2, Y:2 }
  board.placeChecker(2, 2, &checker)
  c := board.takeChecker(2, 2)
  if c == nil {
    t.Error("Cell at ", point.String(), " was empty. Couldn't take checker.")
  }
}

func TestGetChecker(t *testing.T) {
  board := NewBoard(8)
  checker := newChecker(2, 4, false)
  point := point.Point { X:2, Y:4 }
  board.placeChecker(2, 4, &checker)
  c := board.GetChecker(2, 4)
  if c == nil {
    t.Error("Checker at ", point.String(), " is abcent.")
  }
  c = board.GetChecker(5, 5)
  point.X, point.Y = 5, 5
  if c != nil {
    t.Error("Checker at ", point.String(), " should be abcent.")
  }
}

func TestMoveChecker(t *testing.T) {
  board := NewBoard(8)
  point_from := point.Point{ X:3, Y:4 }
  point_to := point.Point{ X:4, Y:3 }
  board.moveChecker(point_from, point_to)
  c := board.GetChecker(3, 4)
  if c != nil {
    t.Error("Checker wasn't moved from ", point_from.String(), " to ", point_to.String())
  }
}
