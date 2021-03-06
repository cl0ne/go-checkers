package checkers

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
	return p.X == q.X && p.Y == q.Y
}

func (p Point) Less(q Point) bool {
	return p.X < q.X || (p.Y < q.Y && p.X == q.X)
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

func (p *Point) Scale(factor int) *Point {
	p.X *= factor
	p.Y *= factor
	return p
}

func (p Point) Scaled(factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}
