package vector

import "math"

var epsilon = math.Nextafter(1.0, 2.0) - 1.0

func FloatEquals(l, r float64) bool {
	return ((l-r) < epsilon && (r-l) < epsilon)
}
