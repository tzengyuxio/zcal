package zcal

import (
	"math"

	"github.com/soniakeys/meeus/base"
	mp "github.com/soniakeys/meeus/moonphase"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/solstice"
)

var secondsOfDay float64 = 24 * 60 * 60

func JDToStemBranch(jd float64) string {
	d := int(math.Floor(jd - JDOfShuodanDongzhi))
	return StemBranch(d)
}

func JDToGanzhi(jd float64) string {
	d := int(math.Floor(jd+.5)) - 11
	return StemBranch(d)
}

func JDToWeekday(jd float64) int {
	d := int(math.Floor(jd+.5)) + 1
	wd := (d % 7)
	if wd < 0 {
		wd += 7
	}
	return wd
}

func CalcDongzhiAndShuo(e *pp.V87Planet, year int, tz float64) (sJD, mJD float64) {
	offset := tz / 24
	sJD = solstice.December2(year, e)
	jy := base.JDEToJulianYear(sJD)
	mJD = mp.New(jy)
	sJD += offset
	mJD += offset
	return
}
