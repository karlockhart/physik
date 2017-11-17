package vector

import "math"

type Vec2 struct {
	X float64
	Y float64
}

func (l Vec2) Equals(r Vec2) bool {
	return FloatEquals(l.X, r.X) && FloatEquals(l.Y, r.Y)
}

func (l Vec2) Add(r Vec2) Vec2 {
	return Vec2{
		X: l.X + r.X,
		Y: l.Y + r.Y,
	}
}

func (l Vec2) Sub(r Vec2) Vec2 {
	return Vec2{
		X: l.X - r.X,
		Y: l.Y - r.Y,
	}
}

func (l Vec2) Mul(r Vec2) Vec2 {
	return Vec2{
		X: l.X * r.X,
		Y: l.Y * r.Y,
	}
}

func (l *Vec2) ScalarMul(r float64) Vec2 {
	return Vec2{
		X: l.X * r,
		Y: l.Y * r,
	}
}

func (l *Vec2) Dot(r Vec2) float64 {
	return l.X*r.X + l.Y*r.Y
}

func (v *Vec2) Magnitude() float64 {
	return math.Sqrt(v.Dot(*v))
}

func (v *Vec2) MagnitudeSq() float64 {
	return v.Dot(*v)
}
