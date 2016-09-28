package zcal

import "math"

var stems = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
var branches = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

// JDOfGongheFirstDay 為西曆前 841 年，共和元年立春日
var JDOfGongheFirstDay = 1414289.5

// JDOfShuodanDongzhi 為西曆 1384年，洪武十七年，朔旦冬至甲子日
var JDOfShuodanDongzhi = 2226910.5

var jdOfGregorianCalendar = 2299160.5

func depart(n float64) (int, float64) {
	i := math.Floor(n)
	f := n - i
	return int(i), f
}

// StemBranch converts a number into Chinese sexagenary cycle representation.
func StemBranch(n int) string {
	g := n % 10
	if g < 0 {
		g += 10
	}
	z := n % 12
	if z < 0 {
		z += 12
	}
	return stems[g] + branches[z]
}

// WesternYearToStemBranch calculates the corresponding stem-branch with the
// given year.
//
// Notice: there is no "0 year" in Gregorian year, the previous year of AD 1 is
// BC 1 (n == -1). If n == 0, the return value will be the same as n == -1.
func WesternYearToStemBranch(n int) string {
	if n < 0 {
		return StemBranch((n - 3) % 60)
	}
	return StemBranch((n - 4) % 60)
}

// WesternCalendarToStemBranch calculates the corresponding stem-branch with
// the given western year, month and day of month
func WesternCalendarToStemBranch(y, m, d int) string {
	var jd float64
	if y > 1582 || (y == 1582 && m >= 10) {
		jd = GregorianCalendarToJD(y, m, d)
	} else {
		if y < 0 {
			y++
		}
		jd = JulianCalendarToJD(y, m, d)
	}
	days := int(jd - JDOfShuodanDongzhi)
	return StemBranch(days)
}

// GregorianCalendarToJD converts Gregorian calendar date to Julian date.
func GregorianCalendarToJD(year, month, day int) float64 {
	a := (14 - month) / 12
	y := year + 4800 - a
	m := month + 12*a - 3
	jdn := day + ((153*m + 2) / 5) + 365*y + (y / 4) - (y / 100) + (y / 400) - 32045

	return float64(jdn) - .5
}

// JDToGregorianCalendar converts Julian date to Gregorian calendar date.
func JDToGregorianCalendar(jd float64) (y, m, d int, t float64) {
	jdn := int(jd + .5)
	f := jdn + 1401 + (((4*jdn+274277)/146097)*3)/4 - 38
	e := 4*f + 3
	g := (e % 1461) / 4
	h := 5*g + 2
	d = (h%153)/5 + 1
	m = (h/153+2)%12 + 1
	y = e/1461 - 4716 + (12+2-m)/12
	t = (jd + .5) - float64(jdn)
	return
}

// JulianCalendarToJD converts Julian calendar date to Julian date.
func JulianCalendarToJD(year, month, day int) float64 {
	a := (14 - month) / 12
	y := year + 4800 - a
	m := month + 12*a - 3
	jdn := day + ((153*m + 2) / 5) + 365*y + (y / 4) - 32083

	return float64(jdn) - .5
}

// JDToJulianCalendar converts Julian date to Julian caldendar date.
func JDToJulianCalendar(jd float64) (y, m, d int, t float64) {
	jdn := int(jd + .5)
	f := jdn + 1401
	e := 4*f + 3
	g := (e % 1461) / 4
	h := 5*g + 2
	d = (h%153)/5 + 1
	m = (h/153+2)%12 + 1
	y = e/1461 - 4716 + (12+2-m)/12
	t = (jd + .5) - float64(jdn)
	return
}

// LeapYearGonghe returns true if year y in the Gonghe year is a leap year
func LeapYearGonghe(y int) bool {
	return (y%4 == 0 && y%100 != 0) || y%500 == 0
}

// JDToGongheCalendar converts Julian date to Gonghe calendar date.
func JDToGongheCalendar(jd float64) (y, m, d int, t float64) {
	gdn, t := depart(jd - JDOfGongheFirstDay)
	α := 0
	if gdn < 0 {
		k := (-gdn / 182621) + 1
		gdn += k * 182621
		α = k * 500
	}
	y = gdn * 500 / 182621
	d = gdn - y*365 - y/4 + y/100 - y/500
	if d < 0 {
		if LeapYearGonghe(y) {
			d += 366
		} else {
			d += 365
		}
		y--
	}
	if LeapYearGonghe(y+1) && d >= 366 {
		d -= 366
		y++
	} else if !LeapYearGonghe(y+1) && d >= 365 {
		d -= 365
		y++
	}
	m, d = d/61, d%61
	if d < 30 {
		m *= 2
	} else {
		m = m*2 + 1
		d -= 30
	}
	y, m, d = y+1-α, m+1, d+1
	return
}

// GongheCalendarToJD converts Gonghe calendar date to Julian date.
func GongheCalendarToJD(year, month, day int) float64 {
	y, m, d := year-1, month-1, day-1
	gdn := y*365 + y/4 - y/100 + y/500
	gdn += m*30 + m/2
	gdn += d
	if y < 0 {
		gdn--
	}

	return float64(gdn) + JDOfGongheFirstDay
}

// GongheCalendarToWesternCalendar converts Gonghe calendar date go Western
// calendar date.
func GongheCalendarToWesternCalendar(y, m, d int) (year, month, day int) {
	jd := GongheCalendarToJD(y, m, d)
	if jd >= jdOfGregorianCalendar {
		year, month, day, _ = JDToGregorianCalendar(jd)
	} else {
		year, month, day, _ = JDToJulianCalendar(jd)
		if year <= 0 {
			year--
		}
	}
	return
}

// WesternCalendarToGongheCalendar converts Western calendar date to Gonghe
// calendar date.
func WesternCalendarToGongheCalendar(y, m, d int) (year, month, day int) {
	var jd float64
	if y > 1582 || (y == 1582 && m >= 10) {
		jd = GregorianCalendarToJD(y, m, d)
	} else {
		if y < 0 {
			y++
		}
		jd = JulianCalendarToJD(y, m, d)
	}
	year, month, day, _ = JDToGongheCalendar(jd)
	return
}
