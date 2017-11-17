package vector

import "math"

type Vec2 struct {
	X float64
	Y float64
}

// Equals returns true if vectors are equal
func (l Vec2) Equals(r Vec2) bool {
	return FloatEquals(l.X, r.X) && FloatEquals(l.Y, r.Y)
}

// Add returns the sum of vectors
func (l Vec2) Add(r Vec2) Vec2 {
	return Vec2{
		X: l.X + r.X,
		Y: l.Y + r.Y,
	}
}

// Sub returns the difference of vectors
func (l Vec2) Sub(r Vec2) Vec2 {
	return Vec2{
		X: l.X - r.X,
		Y: l.Y - r.Y,
	}
}

// Mul returns the product of vectors
func (l Vec2) Mul(r Vec2) Vec2 {
	return Vec2{
		X: l.X * r.X,
		Y: l.Y * r.Y,
	}
}

// ScalarMul returns the product of the vector with a scalar
func (l *Vec2) ScalarMul(r float64) Vec2 {
	return Vec2{
		X: l.X * r,
		Y: l.Y * r,
	}
}

// Dot returns the dot product of vectors
func (l *Vec2) Dot(r Vec2) float64 {
	return l.X*r.X + l.Y*r.Y
}

// Magnitude returns the magnitude of a vector
func (v *Vec2) Magnitude() float64 {
	return math.Sqrt(v.Dot(*v))
}

// MagnitudeSq returns the squared magnitude of a vector
func (v *Vec2) MagnitudeSq() float64 {
	return v.Dot(*v)
}
