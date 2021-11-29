package common

// Point is a point
type Point struct {
	X int
	Y int
}

// Manhattan calculates the Manhattan length for a point
func (p Point) Manhattan() int {
	if p.X < 0 {
		p.X *= -1
	}
	if p.Y < 0 {
		p.Y *= -1
	}
	return p.X + p.Y
}

// ManhattanFrom calculates the Manhattan length for a point from another point
func (p Point) ManhattanFrom(o Point) int {
	p.X -= o.X
	p.Y -= o.Y
	if p.X < 0 {
		p.X *= -1
	}
	if p.Y < 0 {
		p.Y *= -1
	}
	return p.X + p.Y
}
