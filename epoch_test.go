package zcal_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus/solar"

	"github.com/soniakeys/meeus/julian"
	pp "github.com/soniakeys/meeus/planetposition"
	"github.com/soniakeys/meeus/solstice"
	sexa "github.com/soniakeys/sexagesimal"
	"github.com/soniakeys/unit"
	"github.com/stretchr/testify/assert"
	. "github.com/tzengyuxio/zcal"
)

func depart(n float64) (int, float64) {
	i := math.Floor(n) // integer part
	f := n - i         // fractional part
	return int(i), f
}

func jdeSameDay(d1, d2 float64) bool {
	i1, _ := depart(d1 - 0.5)
	i2, _ := depart(d2 - 0.5)
	return (i1 == i2)
}

func TestJDToWeekday(t *testing.T) {
	for _, pair := range []struct {
		jd      float64
		weekday int
	}{
		{-1.75, 6},
		{-1.5, 0},
		{-0.75, 0},
		{-0.5, 1},
		{-0.25, 1},
		{0.5, 2},
		{1.5, 3},
		{2451871.49, 3},
		{2451871.5, 4},
		{2451871.75, 4},
		{2451872.25, 4},
	} {
		wd := JDToWeekday(pair.jd)
		assert.Equal(t, pair.weekday, wd)
	}
}

func TestJDToStemBranch(t *testing.T) {
	for _, pair := range []struct {
		jd         float64
		stemBranch string
	}{
		{-360469.5, "甲子"}, // TZD, The Zero Day
		{-61.5, "壬子"},
		{-60.88, "壬子"},
		{-1.5, "壬子"},
		{-0.88, "壬子"},
		{-0.5, "癸丑"},
		{-0.22, "癸丑"},
		{0.0, "癸丑"},
		{0.11, "癸丑"},
		{0.5, "甲寅"},
		{10.5, "甲子"},
		{15.5, "己巳"},
		{1458495.5, "己巳"},  // -720-02-22
		{1644659.5, "癸丑"},  // -211-11-01
		{1644659.88, "癸丑"}, // -211-11-01
		{1644660.11, "癸丑"}, // -211-11-01
		{2419402.5, "丙子"},  // 1912-01-01
		{2226910.5, "甲子"},  // 1384-12-13, 洪武十七年，朔旦冬至甲子
		{2226910.88, "甲子"}, // 1384-12-13, 洪武十七年，朔旦冬至甲子
		{2226911.11, "甲子"}, // 1384-12-13, 洪武十七年，朔旦冬至甲子
		{2457979.5, "癸酉"},  // 2017-08-14
	} {
		stemBranch := JDToStemBranch(pair.jd)
		assert.Equal(t, pair.stemBranch, stemBranch, "(SteBra) For JD %f expected %s got %s", pair.jd, pair.stemBranch, stemBranch)

		ganzhi := JDToGanzhi(pair.jd)
		assert.Equal(t, pair.stemBranch, ganzhi, "(Ganzhi) For JD %f expected %s got %s", pair.jd, pair.stemBranch, ganzhi)
	}
}

func NoTestFindEpoch(t *testing.T) {
	e, err := pp.LoadPlanetPath(pp.Earth, "/Users/tzengyuxio/SDK/VI_81")
	assert.Nil(t, err)

	for i := 20202; i > -20202; i-- {
		sj, mj := CalcDongzhiAndShuo(e, i, 8)

		sy, sm, sdt := julian.JDToCalendar(sj)
		sd, st := depart(sdt)
		my, mm, mdt := julian.JDToCalendar(mj)
		md, mt := depart(mdt)

		if sy != my || sm != mm || sd != md {
			continue
		}
		gz := JDToStemBranch(sj)
		// gz := JDToGanzhi(sj)
		if gz != "甲子" {
			continue
		}
		fmt.Printf("%5d-%2d-%2d %s | 冬至 %02s 朔日 %02s | Δ: %02s | JDE %14.6f, %14.6f\n",
			sy, sm, sd, gz,
			sexa.FmtTime(unit.TimeFromDay(st)),
			sexa.FmtTime(unit.TimeFromDay(mt)),
			sexa.FmtTime(unit.TimeFromDay(math.Abs(sj-mj))),
			sj, mj)
	}
}

func NoTestFindTZEpoch(t *testing.T) {
	e, err := pp.LoadPlanetPath(pp.Earth, "/Users/tzengyuxio/SDK/VI_81")
	assert.Nil(t, err)

	tz := 8.0 //(110.0 / 15.0)
	offset := tz / 24.0

	// for i := 20202; i > -20202; i-- {
	for i := 10101; i > -10101; i-- {
		sj, mj := CalcDongzhiAndShuo(e, i, 0)

		if !jdeSameDay(sj, mj) {
			continue
		}

		sjPlus := sj + offset
		mjPlus := mj + offset
		if !jdeSameDay(sjPlus, mjPlus) {
			continue
		}

		delta := math.Abs(sj - mj)
		if delta >= 0.5 {
			continue
		}

		gz := JDToStemBranch(sj)
		if gz != "甲子" {
			continue
		}
		gzPlus := JDToStemBranch(sjPlus)
		if gzPlus != "甲子" {
			continue
		}

		sy, sm, sdt := julian.JDToCalendar(sj)
		sd, st := depart(sdt)
		_, _, mdt := julian.JDToCalendar(mj)
		_, mt := depart(mdt)

		fmt.Printf("%5d-%2d-%2d %s | 冬至 %02s 朔日 %02s | Δ: %02s | JDE %14.6f, %14.6f | WD %d\n",
			sy, sm, sd, gz,
			sexa.FmtTime(unit.TimeFromDay(st)),
			sexa.FmtTime(unit.TimeFromDay(mt)),
			sexa.FmtTime(unit.TimeFromDay(delta)),
			sj, mj,
			JDToWeekday(sj))

		_, _, sdtPlus := julian.JDToCalendar(sjPlus)
		_, stPlus := depart(sdtPlus)
		_, _, mdtPlus := julian.JDToCalendar(mjPlus)
		_, mtPlus := depart(mdtPlus)

		fmt.Printf("      [GMT+%4.2f] | 冬至 %02s 朔日 %02s | Δ: %02s | JDE %14.6f, %14.6f\n",
			tz,
			sexa.FmtTime(unit.TimeFromDay(stPlus)),
			sexa.FmtTime(unit.TimeFromDay(mtPlus)),
			sexa.FmtTime(unit.TimeFromDay(delta)),
			sjPlus, mjPlus)
	}
}

func NoTestStemBranchAndGanzhi(t *testing.T) {
	for i := 1644659.5; i > -1644659.5; i-- {
		sb := JDToStemBranch(i)
		gz := JDToGanzhi(i)
		// fmt.Println(i)
		assert.Equal(t, sb, gz, "wtf? %f", i)
	}
}

func TestFindGHZeroYear(t *testing.T) {
	e, err := pp.LoadPlanetPath(pp.Earth, "/Users/tzengyuxio/SDK/VI_81")
	assert.Nil(t, err)

	jd := julian.CalendarJulianToJD(-842, 2, 12.260417)
	fmt.Printf("zero gonghe %f\n", jd)

	var jdw, jds float64
	for i := -842; i <= -840; i++ {
		jdw = solstice.December2(i-1, e)
		jds = solstice.March2(i, e)
		fmt.Printf("GH Year(%d): %f (%0.2f)\n", i, (jdw+jds)/2, jds-jdw)
		// p, b0, l0 := solardisk.Ephemeris(jds, e)
		// fmt.Printf("P, B0, L0: %.2f, %+2.f, %.2f\n", p, b0, l0)
		alpha, delta, _ := solar.ApparentEquatorialVSOP87(e, jds)
		fmt.Printf("α: %.3d  δ: %+.2d\n", sexa.FmtRA(alpha), sexa.FmtAngle(delta))
	}

	fmt.Println("--------------------")
	fmt.Printf("GH Year(%d): %f\n", -842, jd)
	alpha, delta, _ := solar.ApparentEquatorialVSOP87(e, jd)
	fmt.Printf("JD: %08.6f, α: %.3d  δ: %+.2d\n", jd, sexa.FmtRA(alpha), sexa.FmtAngle(delta))
	jdplus := jd - 2.866585
	alpha, delta, _ = solar.ApparentEquatorialVSOP87(e, jdplus)
	fmt.Printf("JD: %08.6f, α: %.3d  δ: %+.2d\n", jdplus, sexa.FmtRA(alpha), sexa.FmtAngle(delta))
	// fmt.Printf("JD: %08.6f, α: %.3d  δ: %+.2d\n", jdplus, sexa.FmtAngle(unit.RAFromRad(alpha)), sexa.FmtAngle(unit.AngleFromDeg(delta)))
	fmt.Println("--------------------")
	lambda, beta, _ := solar.ApparentVSOP87(e, jd)
	fmt.Printf("JD: %08.6f, λ: %.3d  β: %+.2d\n", jd, sexa.FmtAngle(lambda), sexa.FmtAngle(beta)) // 查書的春分點
	jdplus2 := jd - 0.304905
	lambda, beta, _ = solar.ApparentVSOP87(e, jdplus2)
	fmt.Printf("JD: %08.6f, λ: %.3d  β: %+.2d\n", jdplus2, sexa.FmtAngle(lambda), sexa.FmtAngle(beta)) // 計算出來的春分點
	fmt.Println("--------------------")

	jd2 := JulianCalendarToJD(-841, 2, 12)
	y, m, d, _ := JDToGHC(jd2) // (1414289.5) // -841-02-12
	fmt.Printf("GH Year 1(old) is %02d-%02d-%02d (%f)\n", y, m, d, jd2)
}
