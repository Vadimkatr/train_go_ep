package figure

import (
	"math"
	"testing"
)

func TestFigureArea(t *testing.T) {
	var tests = []struct {
		fg        Figure
		area      float64
		perimeter float64
	}{
		{
			Square{
				Start: Point{X: 1, Y: 1}, A: 4, B: 5,
			}, 20, 18,
		},
		{
			Square{
				Start: Point{X: -2, Y: 2}, A: 1, B: 1,
			}, 1, 4,
		},
		{
			Square{
				Start: Point{X: 11, Y: 15}, A: 0, B: 1,
			}, 0, 0,
		},
		{
			Square{
				Start: Point{X: 11, Y: 15}, A: 10, B: 20,
			}, 200, 60,
		},
		{
			Circle{
				Center: Point{X: 1, Y: 1}, R: 4,
			}, math.Pi * 4 * 4, 2 * math.Pi * 4,
		},
		{
			Circle{
				Center: Point{X: -2, Y: 2}, R: 1,
			}, math.Pi * 1 * 1, 2 * math.Pi * 1,
		},
		{
			Circle{
				Center: Point{X: 11, Y: 15}, R: 0,
			}, math.Pi * 0 * 0, 2 * math.Pi * 0,
		},
	}
	for i, c := range tests {
		gotArea := c.fg.Area()
		if gotArea != c.area {
			t.Errorf("Interface Figure{}.Area(%v) == %f, but must be %f; (watch test index %d for  details).",
				c.fg, gotArea, c.area, i)
		}

		gotPerim := c.fg.Perimeter()
		if gotPerim != c.perimeter {
			t.Errorf("Interface Figure{}.Perimeter(%v) == %f, but must be %f; (watch test index %d for  details).",
				c.fg, gotPerim, c.perimeter, i)
		}
	}
}
