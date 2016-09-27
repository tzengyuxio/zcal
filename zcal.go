package zcal

var stems = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
var branches = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

// JDOfGongheFirstDay 為西曆前 841 年，共和元年立春日
var JDOfGongheFirstDay = 1414289.5

// JDOfShuodanDongzhi 為西曆 1384年，洪武十七年，朔旦冬至甲子日
var JDOfShuodanDongzhi = 2226910.5

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
	return 0, 0, 0, 0.0
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
	return 0, 0, 0, 0.0
}
