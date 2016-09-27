package zcal

var stems = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
var branches = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

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

// GregorianYearToStemBranch calculates the corresponding stem-branch with the given year.
func GregorianYearToStemBranch(n int) string {
	if n < 0 {
		return StemBranch((n - 3) % 60)
	}
	return StemBranch((n - 4) % 60)
}
