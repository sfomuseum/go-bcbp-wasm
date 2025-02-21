// Core mathematical functions for astronomical calculations.
package mathutils

import "math"

const _DEG_RAD = math.Pi / 180
const _RAD_DEG = 180 / math.Pi
const _PI2 = math.Pi * 2

// Compares two floats with a given precision.
func AlmostEqual(a, b, threshold float64) bool {
	return math.Abs(a-b) <= threshold
}

// Fractional part of a number.
//
// It uses the standard [math.Modf] function.
// The result always keeps sign of the argument.
//
//	Frac(-5.5) = -0.5
func Frac(x float64) float64 {
	_, f := math.Modf(x)
	return f
}

func ToRange(x float64, r float64) float64 {
	z := math.Mod(x, r)
	if z < 0 {
		return z + r
	}
	return z
}

// Reduces hours to range 0 >= x < 24
func ReduceHours(hours float64) float64 {
	return ToRange(hours, 24)
}

// Reduces arc-degrees to range 0 >= x < 360
func ReduceDeg(deg float64) float64 {
	return ToRange(deg, 360)
}

// Reduces radians to range 0 >= x < 2 * pi
func ReduceRad(rad float64) float64 {
	return ToRange(rad, _PI2)
}

// Calculates polynome:
//
//	a1 + a2*t + a3*t*t + a4*t*t*t...
//
// t is a number of Julian centuries elapsed since 1900, Jan 0.5
// terms is a list of terms
func Polynome(t float64, terms ...float64) float64 {
	res := 0.0
	for i, k := range terms {
		p := math.Pow(t, float64(i))
		res += k * p
	}
	return res
}

// Converts arc-degrees to radians
func Radians(deg float64) float64 {
	return deg * _DEG_RAD
}

// Converts radians to arc-degrees
func Degrees(rad float64) float64 {
	return rad * _RAD_DEG
}

// Used with polinomial function for better accuracy.
func Frac360(x float64) float64 {
	return Frac(x) * 360
}

// Converts decimal hours to sexagesimal values.
//
// If x is < 0, then the first non-zero return value will be negative.
//
//	Hms(-0.5) = 0, -30, 0.0
func Hms(x float64) (hours int, minutes int, seconds float64) {
	i, f := math.Modf(math.Abs(x))
	hours = int(i)
	i, f = math.Modf(f * 60)
	minutes = int(i)
	seconds = f * 60
	if x < 0 {
		if hours != 0 {
			hours = -hours
		} else if minutes != 0 {
			minutes = -minutes
		} else if seconds != 0 {
			seconds = -seconds
		}
	}
	return hours, minutes, seconds
}
