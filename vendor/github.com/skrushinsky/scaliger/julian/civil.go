package julian

import (
	"math"

	"github.com/skrushinsky/scaliger/mathutils"
)

// Civil calendar date, usually Gregorian.
// Time is represented by fractional part of a day.
// For example,
//
//	7h30m UT is (7 + 30 / 60) / 24 = 0.3125.
type CivilDate struct {
	// year, astronomical, negative for BC dates
	Year int
	// month number, 1-12
	Month int
	// day, fractional part represents hours
	Day float64
}

// Compares two dates.
func EqualDates(a, b CivilDate) bool {
	return a.Year == b.Year && a.Month == b.Month && mathutils.AlmostEqual(a.Day, b.Day, 1e-6)
}

// Returns true ig given year is a leap year.
func IsLeapYear(year int) bool {
	return (year%4 == 0) && ((year%100 != 0) || (year%400 == 0))
}

// Number of days in the year up to a particular date.
func DayOfYear(date CivilDate) int {
	var k float64
	if IsLeapYear(date.Year) {
		k = 1
	} else {
		k = 2
	}
	mo := float64(date.Month)
	a := math.Floor(275 * mo / 9)
	b := math.Floor(k * ((mo + 9) / 12))
	c := math.Floor(date.Day)
	return int(a-b+c) - 30
}
