package point

import (
  "fmt"
  "math"
)

type Point struct {
  x, y int
}

func (p Point) Add(q Point) Point {
  return Point {p.x + q.x, p.y + q.y}
}

func (p Point) Sub(q Point) Point {
  return Point {p.x - q.x, p.y - q.y}
}

func (p Point) Equal(q Point) bool {
  if p.x == q.x && p.y == q.y {
    return true
  }
  return false
}

func (p Point) String() string {
  return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func (p Point) Manhattan() int {
  return int(math.Abs(float64(0 - p.x)) + math.Abs(float64(0 - p.y)))
}

func (p Point) ManhattanTo(q Point) int {
  return int(math.Abs(float64(p.x - q.x)) + math.Abs(float64(p.y - q.y)))
}
