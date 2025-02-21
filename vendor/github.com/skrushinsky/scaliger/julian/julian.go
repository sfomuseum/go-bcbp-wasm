// The main purpose is to convert between civil dates and Julian dates.
// Julian date (JD) is the number of days elapsed since mean UT noon of
// January 1st 4713 BC. This system of time measurement is widely adopted by
// the astronomers.
package julian

import (
	"math"
	"time"

	"github.com/skrushinsky/scaliger/mathutils"
)

// Seconds per day
const SEC_PER_DAY = 24 * 60 * 60

// Number of days in a second
const DAYS_PER_SEC = 1.0 / SEC_PER_DAY

// Days per century
const DAYS_PER_CENT = 36525

// Julian day for 2000 Jan. 1.5
const J2000 = 2451545.0

// Julian day for 1900 Jan. 0.5
const J1900 = 2415020.0

// Dates before Oct 10, 1582 are considered non-gregorian
func isGregorian(date CivilDate) bool {
	if date.Year > 1582 {
		return true
	}
	if date.Year < 1582 {
		return false
	}
	if date.Month > 10 {
		return true
	}
	if date.Month < 10 {
		return false
	}
	if date.Day > 10 {
		return true
	}
	return false
}

// Converts calendar date into Julian days.
func CivilToJulian(date CivilDate) float64 {
	var y, m float64
	if date.Month > 2 {
		y = float64(date.Year)
		m = float64(date.Month)
	} else {
		y = float64(date.Year) - 1
		m = float64(date.Month) + 12
	}

	var t float64
	if date.Year < 0 {
		t = 0.75
	}

	var a, b float64
	if isGregorian(date) {
		a = math.Trunc(y / 100)
		b = 2 - a + math.Trunc(a/4)
	}

	return b + math.Trunc(365.25*y-t) + math.Trunc(30.6001*(m+1)) + date.Day + 1720994.5
}

// Converts number of Julian days into the calendar date.
func JulianToCivil(jd float64) CivilDate {
	i, f := math.Modf(jd + 0.5)

	var b float64
	if i > 2299160 {
		a := math.Trunc((i - 1867216.25) / 36524.25)
		b = i + 1 + a - math.Trunc(a/4)
	} else {
		b = i
	}
	c := b + 1524
	d := math.Trunc((c - 122.1) / 365.25)
	e := math.Trunc(365.25 * d)
	g := math.Trunc((c - e) / 30.6001)

	da := c - e + f - math.Trunc(30.6001*g)
	var mo float64
	if g < 13.5 {
		mo = g - 1
	} else {
		mo = g - 13
	}
	var ye float64
	if mo > 2.5 {
		ye = d - 4716
	} else {
		ye = d - 4715
	}

	return CivilDate{Year: int(ye), Month: int(mo), Day: da}
}

// Given number of Julian days, calculates JD at Greenwich midnight.
func JulianMidnight(jd float64) float64 {
	return math.Floor(jd-0.5) + 0.5
}

// Julian Day corresponding to January 0.0 of a given year.
//
// Zero day is a special case of date: it indicates 12h UT of previous calendar
// date. For instance, *1900 January 0.5* is often used instead of
// *1899 December 31.5* to designate start of the astronomical epoch.
func JulianDateZero(year int) float64 {
	y := float64(year - 1)
	a := math.Trunc(y / 100)
	return math.Trunc(365.25*y) - a + math.Trunc(a/4) + 1721424.5
}

// Converts fractional part of a Julian Date to UTC as decimal hours.
func ExtractUTC(jd float64) float64 {
	return (jd - JulianMidnight(jd)) * 24
}

// Given an date string, calculate Julian Date.
//
// The date must be in RFC3339 format without time zone offset, i.e.:
//
//	jd, _ := DateStringToJulian("2006-01-02T15:04:05Z")
func DateStringToJulian(date string) (float64, error) {
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return 0, err
	}
	ut := float64(t.Hour()) + float64(t.Minute())/60
	jd := CivilToJulian(CivilDate{Year: t.Year(), Month: int(t.Month()), Day: float64(t.Day()) + ut/24})
	return jd, nil
}

// Given Julian Date return RFC-3339 formatted date string.
func JulianToDateString(jd float64) string {
	civil := JulianToCivil(jd)

	var year int
	if civil.Year < 1 {
		// for dates BC convert astronomical year to civil
		year = civil.Year - 1
	} else {
		year = civil.Year
	}
	i, f := math.Modf(civil.Day)
	day := int(i)
	hour, min, sec := mathutils.Hms(f * 24)
	nanos := int(mathutils.Frac(sec) * 1e9)
	date := time.Date(year, time.Month(civil.Month), day, hour, min, int(sec), nanos, time.UTC)
	return date.Format(time.RFC3339)
}
