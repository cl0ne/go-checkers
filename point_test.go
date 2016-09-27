package point

import "testing"

func TestEqual(t *testing.T) {
	cases := []struct {
		a, b   Point
		result bool
	}{
		{Point{x: -1, y: 1}, Point{x: -1, y: 1}, true},
		{Point{x: -1, y: 1}, Point{x: 0, y: 0}, false},
		{Point{}, Point{}, true},
	}
	for _, c := range cases {
		if c.a.Equal(c.b) != c.result {
			t.Error(c.a, " == ", c.b, " must be ", c.result)
		} else if c.b.Equal(c.a) != c.result {
			t.Error(c.a, " == ", c.b, " is not commutative")
		}
	}
}

func TestAdd(t *testing.T) {
	cases := []struct {
		a, b, result Point
	}{
		{Point{x: 1, y: 2}, Point{x: -1, y: 1}, Point{x: 0, y: 3}},
		{Point{}, Point{x: 0, y: 0}, Point{}},
		{Point{x: 1, y: 1}, Point{}, Point{x: 1, y: 1}},
		{Point{x: 1, y: 1}, Point{x: 1, y: 1}, Point{x: 2, y: 2}},
	}
	for _, c := range cases {
		r := c.a.Add(c.b)
		if r != c.result {
			t.Error(c.a, " + ", c.b, " expected ", c.result, ", got ", r)
		} else if r := c.b.Add(c.a); r != c.result {
			t.Error(c.a, " + ", c.b, " is not commutative: expected ", c.result, ", got ", r)
		}
	}
}

func TestSub(t *testing.T) {
	cases := []struct {
		a, b, result Point
	}{
		{Point{x: 1, y: 2}, Point{x: -1, y: 1}, Point{x: 2, y: 1}},
		{Point{}, Point{x: 0, y: 0}, Point{}},
		{Point{x: 1, y: 1}, Point{}, Point{x: 1, y: 1}},
		{Point{x: 1, y: 1}, Point{x: 1, y: 1}, Point{x: 0, y: 0}},
	}
	for _, c := range cases {
		r := c.a.Sub(c.b)
		if r != c.result {
			t.Error(c.a, " - ", c.b, " expected ", c.result, ", got ", r)
		}
	}
}

func TestString(t *testing.T) {
	cases := []struct {
		p Point
		s string
	}{
		{Point{x: 1, y: 2}, "(1,2)"},
		{Point{}, "(0,0)"},
		{Point{x: 1, y: 1}, "(1,1)"},
		{Point{x: 1, y: -1}, "(1,-1)"},
		{Point{x: 0, y: -5}, "(0,-5)"},
	}
	for _, c := range cases {
		s := c.p.String()
		if s != c.s {
			t.Error("Point.String: expected ", c.s, ", got ", s)
		}
	}
}

func TestManhattan(t *testing.T) {
	cases := []struct {
		p Point
		d int
	}{
		{Point{x: 1, y: 2}, 3},
		{Point{}, 0},
		{Point{x: 1, y: -1}, 2},
		{Point{x: 0, y: -5}, 5},
	}
	for _, c := range cases {
		d := c.p.Manhattan()
		if d != c.d {
			t.Error(c.p, " Manhattan length expected ", c.d, ", got ", d)
		}
	}
}

func TestManhattanTo(t *testing.T) {
	cases := []struct {
		a, b Point
		d    int
	}{
		{Point{}, Point{x: 1, y: 2}, 3},
		{Point{}, Point{}, 0},
		{Point{}, Point{x: 1, y: -1}, 2},
		{Point{}, Point{x: 0, y: -5}, 5},
		{Point{}, Point{x: -2, y: -5}, 7},

		{Point{x: 1, y: -1}, Point{x: 1, y: 2}, 3},
		{Point{x: 1, y: 2}, Point{x: 1, y: -1}, 3},
		{Point{x: -1, y: -2}, Point{x: -3, y: -4}, 4},
		{Point{x: 1, y: 2}, Point{x: 3, y: 4}, 4},
	}
	for _, c := range cases {
		d := c.a.ManhattanTo(c.b)
		if d != c.d {
			t.Error(c.a, " Manhattan distance to, ", c.b, " expected ", c.d, ", got ", d)
		}
	}
}
