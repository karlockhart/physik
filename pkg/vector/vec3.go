package vector

import "math"

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

// Equals returns true if vectors are equal
func (l *Vec3) Equals(r Vec3) bool {
	return FloatEquals(l.X, r.X) && FloatEquals(l.Y, r.Y) && FloatEquals(l.Z, r.Z)
}

// Add returns the sum of vectors
func (l *Vec3) Add(r Vec3) Vec3 {
	return Vec3{
		X: l.X + r.X,
		Y: l.Y + r.Y,
		Z: l.Z + r.Z,
	}
}

// Sub returns the difference of vectors
func (l *Vec3) Sub(r Vec3) Vec3 {
	return Vec3{
		X: l.X - r.X,
		Y: l.Y - r.Y,
		Z: l.Z - r.Z,
	}
}

// Mul returns the product of vectors
func (l *Vec3) Mul(r Vec3) Vec3 {
	return Vec3{
		X: l.X * r.X,
		Y: l.Y * r.Y,
		Z: l.Z * r.Z,
	}
}

// ScalarMul returns the product of the vector with a scalar
func (l *Vec3) ScalarMul(r float64) Vec3 {
	return Vec3{
		X: l.X * r,
		Y: l.Y * r,
		Z: l.Z * r,
	}
}

// Dot returns the dot product of vectors
func (l *Vec3) Dot(r Vec3) float64 {
	return l.X*r.X + l.Y*r.Y + l.Z*r.Z
}

// Magnitude returns the magnitude of a vector
func (v *Vec3) Magnitude() float64 {
	return math.Sqrt(v.Dot(*v))
}

// MagnitudeSq returns the squared magnitude of a vector
func (v *Vec3) MagnitudeSq() float64 {
	return v.Dot(*v)
}
