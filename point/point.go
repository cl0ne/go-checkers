package point

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func (p Point) Equal(q Point) bool {
	if p.X == q.X && p.Y == q.Y {
		return true
	}
	return false
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func (p Point) Manhattan() int {
	return int(math.Abs(float64(0-p.X)) + math.Abs(float64(0-p.Y)))
}

func (p Point) ManhattanTo(q Point) int {
	return int(math.Abs(float64(p.X-q.X)) + math.Abs(float64(p.Y-q.Y)))
}
