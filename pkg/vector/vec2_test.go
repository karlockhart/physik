package vector

import (
	"math"
	"testing"
)

func TestVec2Equals(t *testing.T) {
	l := Vec2{X: 1, Y: 1}
	r := Vec2{X: 1, Y: 1}
	if !l.Equals(r) {
		t.Fail()
	}
}

func TestVec2Add(t *testing.T) {
	l := Vec2{X: 1, Y: 1}
	r := Vec2{X: 1, Y: 1}
	res := l.Add(r)

	if !res.Equals(Vec2{X: 2, Y: 2}) {
		t.Fail()
	}
}

func TestVec2Sub(t *testing.T) {
	l := Vec2{X: 2, Y: 2}
	r := Vec2{X: 1, Y: 1}
	res := l.Sub(r)

	if !res.Equals(Vec2{X: 1, Y: 1}) {
		t.Fail()
	}
}

func TestVec2Mul(t *testing.T) {
	l := Vec2{X: 2, Y: 2}
	r := Vec2{X: 2, Y: 3}
	res := l.Mul(r)

	if !res.Equals(Vec2{X: 4, Y: 6}) {
		t.Fail()
	}
}

func TestVec2ScalarMul(t *testing.T) {
	l := Vec2{X: 2, Y: 3}
	r := 2.0
	res := l.ScalarMul(r)

	if !res.Equals(Vec2{X: 4, Y: 6}) {
		t.Fail()
	}
}

func TestVec2Dot(t *testing.T) {
	l := Vec3{X: 2, Y: 2}
	r := Vec3{X: 2, Y: 3}
	res := l.Dot(r)

	if !FloatEquals(res, 10.0) {
		t.Fail()
	}
}

func TestVec2Magnitude(t *testing.T) {
	l := Vec2{X: 3, Y: 3}
	res := l.Magnitude()

	if !FloatEquals(res, math.Sqrt(18)) {
		t.Fail()
	}
}

func TestVec2MagnitudeSquare(t *testing.T) {
	l := Vec2{X: 3, Y: 3}
	res := l.MagnitudeSq()

	if !FloatEquals(res, 18) {
		t.Fail()
	}
}
