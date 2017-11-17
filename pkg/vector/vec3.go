package vector

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (l *Vec3) Equals(r Vec3) bool {
	return FloatEquals(l.X, r.X) && FloatEquals(l.Y, r.Y) && FloatEquals(l.Z, r.Z)
}

func (l *Vec3) Add(r Vec3) Vec3 {
	return Vec3{
		X: l.X + r.X,
		Y: l.Y + r.Y,
		Z: l.Z + r.Z,
	}
}

func (l *Vec3) Sub(r Vec3) Vec3 {
	return Vec3{
		X: l.X - r.X,
		Y: l.Y - r.Y,
		Z: l.Z - r.Z,
	}
}

func (l *Vec3) Mul(r Vec3) Vec3 {
	return Vec3{
		X: l.X * r.X,
		Y: l.Y * r.Y,
		Z: l.Z * r.Z,
	}
}

func (l *Vec3) ScalarMul(r float64) Vec3 {
	return Vec3{
		X: l.X * r,
		Y: l.Y * r,
		Z: l.Z * r,
	}
}

func (l *Vec3) Dot(r Vec3) float64 {
	return l.X*r.X + l.Y*r.Y + l.Z*r.Z
}
