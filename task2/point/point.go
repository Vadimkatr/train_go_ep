package point

type Point struct {
	X, Y int
}

type Square struct {
	Start Point
	A uint
}

func (s Square) End() Point {
	return Point{
		s.Start.X + int(s.A), 
		s.Start.Y - int(s.A),
	}
}

func (s Square) Perimeter() uint {
	return s.A * 4
}

func (s Square) Area() uint {
	return s.A * s.A
}
