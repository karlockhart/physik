package vector

import (
	"math"
	"testing"
)

func TestVec3Equals(t *testing.T) {
	l := Vec3{X: 1, Y: 1, Z: 1}
	r := Vec3{X: 1, Y: 1, Z: 1}
	if !l.Equals(r) {
		t.Fail()
	}
}

func TestVec3Add(t *testing.T) {
	l := Vec3{X: 1, Y: 1, Z: 1}
	r := Vec3{X: 1, Y: 1, Z: 1}
	res := l.Add(r)

	if !res.Equals(Vec3{X: 2, Y: 2, Z: 2}) {
		t.Fail()
	}
}

func TestVec3Sub(t *testing.T) {
	l := Vec3{X: 2, Y: 2, Z: 2}
	r := Vec3{X: 1, Y: 1, Z: 1}
	res := l.Sub(r)

	if !res.Equals(Vec3{X: 1, Y: 1, Z: 1}) {
		t.Fail()
	}
}

func TestVec3Mul(t *testing.T) {
	l := Vec3{X: 2, Y: 2, Z: 2}
	r := Vec3{X: 2, Y: 3, Z: 4}
	res := l.Mul(r)

	if !res.Equals(Vec3{X: 4, Y: 6, Z: 8}) {
		t.Fail()
	}
}

func TestVec3ScalarMul(t *testing.T) {
	l := Vec3{X: 2, Y: 3, Z: 4}
	r := 2.0
	res := l.ScalarMul(r)

	if !res.Equals(Vec3{X: 4, Y: 6, Z: 8}) {
		t.Fail()
	}
}

func TestVec3Dot(t *testing.T) {
	l := Vec3{X: 2, Y: 2, Z: 2}
	r := Vec3{X: 2, Y: 3, Z: 4}
	res := l.Dot(r)

	if !FloatEquals(res, 18.0) {
		t.Fail()
	}
}

func TestVec3Magnitude(t *testing.T) {
	l := Vec3{X: 3, Y: 3, Z: 3}
	res := l.Magnitude()

	if !FloatEquals(res, math.Sqrt(27)) {
		t.Fail()
	}
}

func TestVec3MagnitudeSquare(t *testing.T) {
	l := Vec3{X: 3, Y: 3, Z: 3}
	res := l.MagnitudeSq()

	if !FloatEquals(res, 27) {
		t.Fail()
	}
}
