package zcal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/tzengyuxio/zcal"
)

func TestWesternYearToStemBranch(t *testing.T) {
	for _, pair := range []struct {
		year       int
		stemBranch string
	}{
		{-2697, "甲子"}, // 黃帝紀元元年
		{-841, "庚申"},  // 共和元年，中國確切紀年之始
		{-221, "庚辰"},  // 秦始皇二十六年，統一中國
		{4, "甲子"},     // 黃帝紀元 45 甲子 (2700 年) 之後
		{1384, "甲子"},  // 洪武十七年
		{1894, "甲午"},  // 甲午戰爭
		{1912, "壬子"},  // 中華民國元年
		{2012, "壬辰"},
	} {
		stemBranch := WesternYearToStemBranch(pair.year)
		assert.Equal(t, pair.stemBranch, stemBranch, "For year %d expected %s got %s", pair.year, pair.stemBranch, stemBranch)
	}
}

func TestWesternCalendarToStemBranch(t *testing.T) {
	for _, pair := range []struct {
		y, m, d    int
		stemBranch string
	}{
		{-841, 2, 12, "癸未"},  // 共和元年立春
		{-720, 2, 22, "己巳"},  // 魯隱公三年夏曆二月己巳日，西元前 720 年 2 月 22 日
		{-211, 11, 1, "癸丑"},  // 三十七年十月癸丑，始皇出遊
		{1384, 12, 13, "甲子"}, // 洪武十七年，朔旦冬至甲子
		{1912, 2, 18, "甲子"},
		{9912, 2, 18, "甲子"},
	} {
		stemBranch := WesternCalendarToStemBranch(pair.y, pair.m, pair.d)
		assert.Equal(t, pair.stemBranch, stemBranch, "For date %04d-%02d-%02d expected %s got %s", pair.y, pair.m, pair.d, pair.stemBranch, stemBranch)
	}
}

func TestGregorianCalendarToJD(t *testing.T) {
	for _, pair := range []struct {
		y, m, d int
		jd      float64
	}{
		{-4699, 1, 1, 4785.5},
		{-2114, 2, 12, 948979.5},
		{-1050, 3, 23, 1337636.5},
		{-123, 4, 14, 1676238.5},
		{-1, 5, 25, 1720838.5},
		{0, 6, 16, 1721226.5},
		{1, 7, 27, 1721632.5},
		{123, 8, 8, 1766203.5},
		{1678, 9, 9, 2334188.5},
		{2000, 1, 1, 2451544.5},
		{2000, 10, 10, 2451827.5},
	} {
		jd := GregorianCalendarToJD(pair.y, pair.m, pair.d)
		assert.Equal(t, pair.jd, jd, "For date %04d-%02d-%02d expected %.1f got %.1f", pair.y, pair.m, pair.d, pair.jd, jd)
	}
}

func TestJulianCalendarToJD(t *testing.T) {
	for _, pair := range []struct {
		y, m, d int
		jd      float64
	}{
		{-4712, 1, 2, 0.5},
		{-4699, 1, 1, 4748.5},
		{-2114, 2, 12, 948961.5},
		{-1050, 3, 23, 1337626.5},
		{-840, 2, 12, JDOfGongheFirstDay},
		{-123, 4, 14, 1676235.5},
		{-1, 5, 25, 1720836.5},
		{0, 6, 16, 1721224.5},
		{1, 7, 27, 1721630.5},
		{123, 8, 8, 1766202.5},
		{1384, 12, 13, JDOfShuodanDongzhi},
		{1678, 10, 9, 2334228.5},
		{2000, 11, 10, 2451871.5},
	} {
		jd := JulianCalendarToJD(pair.y, pair.m, pair.d)
		assert.Equal(t, pair.jd, jd, "For date %04d-%02d-%02d expected %.1f got %.1f", pair.y, pair.m, pair.d, pair.jd, jd)
	}
}

func TestJDToGregorianCalendar(t *testing.T) {
	for _, pair := range []struct {
		y, m, d int
		jd      float64
	}{
		{-4699, 1, 1, 4785.5},
		{-2114, 2, 12, 948979.5},
		{-1050, 3, 23, 1337636.5},
		{-123, 4, 14, 1676238.5},
		{-1, 5, 25, 1720838.5},
		{0, 6, 16, 1721226.5},
		{1, 7, 27, 1721632.5},
		{123, 8, 8, 1766203.5},
		{1678, 9, 9, 2334188.5},
		{2000, 1, 1, 2451544.5},
		{2000, 10, 10, 2451827.5},
	} {
		y, m, d, _ := JDToGregorianCalendar(pair.jd)
		assert.Equal(t, pair.y, y, "For JD %.4f expected year %02d got %02d", pair.jd, pair.y, y)
		assert.Equal(t, pair.m, m, "For JD %.4f expected month %02d got %02d", pair.jd, pair.m, m)
		assert.Equal(t, pair.d, d, "For JD %.4f expected day %02d got %02d", pair.jd, pair.d, d)
	}
}

func TestJDToJulianCalendar(t *testing.T) {
	for _, pair := range []struct {
		y, m, d int
		jd      float64
	}{
		{-4712, 1, 2, 0.5},
		{-4699, 1, 1, 4748.5},
		{-2114, 2, 12, 948961.5},
		{-1050, 3, 23, 1337626.5},
		{-123, 4, 14, 1676235.5},
		{-1, 5, 25, 1720836.5},
		{0, 6, 16, 1721224.5},
		{1, 7, 27, 1721630.5},
		{123, 8, 8, 1766202.5},
		{1384, 12, 13, JDOfShuodanDongzhi},
		{1678, 10, 9, 2334228.5},
		{2000, 11, 10, 2451871.5},
	} {
		y, m, d, _ := JDToJulianCalendar(pair.jd)
		assert.Equal(t, pair.y, y, "For JD %.4f expected year %02d got %02d", pair.jd, pair.y, y)
		assert.Equal(t, pair.m, m, "For JD %.4f expected month %02d got %02d", pair.jd, pair.m, m)
		assert.Equal(t, pair.d, d, "For JD %.4f expected day %02d got %02d", pair.jd, pair.d, d)
	}
}

func TestJDToGongheCalendar(t *testing.T) {
	for _, pair := range []struct {
		y, m, d int
		jd      float64
	}{
		{-4, 4, 4, JDOfGongheFirstDay - 1733.0},
		{-2, 2, 2, JDOfGongheFirstDay - 1065.0},
		{-1, 1, 1, JDOfGongheFirstDay - 731.0},
		{-1, 12, 30, JDOfGongheFirstDay - 367.0},
		{0, 1, 1, JDOfGongheFirstDay - 366.0},
		{0, 1, 2, JDOfGongheFirstDay - 365.0},
		{0, 12, 31, JDOfGongheFirstDay - 1.0},
		{1, 1, 1, JDOfGongheFirstDay},
		{1, 1, 2, JDOfGongheFirstDay + 1.0},
		{1, 1, 30, JDOfGongheFirstDay + 29.0},
		{1, 2, 1, JDOfGongheFirstDay + 30.0},
		{1, 12, 30, JDOfGongheFirstDay + 364.0},
		{2, 2, 2, JDOfGongheFirstDay + 396.0},
		{4, 1, 1, JDOfGongheFirstDay + 1095.0},
		{4, 12, 31, JDOfGongheFirstDay + 1460.0},
		{5, 1, 1, JDOfGongheFirstDay + 1461.0},
		{100, 12, 30, JDOfGongheFirstDay + 36523.0},
		{101, 1, 1, JDOfGongheFirstDay + 36524.0},
		{500, 12, 31, JDOfGongheFirstDay + 182620.0},
		{501, 1, 1, JDOfGongheFirstDay + 182621.0},
		{2225, 11, 19, JDOfShuodanDongzhi},
	} {
		y, m, d, _ := JDToGongheCalendar(pair.jd)
		assert.Equal(t, pair.y, y, "For JD %.4f expected year %02d got %02d", pair.jd, pair.y, y)
		assert.Equal(t, pair.m, m, "For JD %.4f expected month %02d got %02d", pair.jd, pair.m, m)
		assert.Equal(t, pair.d, d, "For JD %.4f expected day %02d got %02d", pair.jd, pair.d, d)
	}
}

func TestGongheCalendarToJD(t *testing.T) {
	for _, pair := range []struct {
		y, m, d int
		jd      float64
	}{
		{-4, 4, 4, JDOfGongheFirstDay - 1733.0},
		{-2, 2, 2, JDOfGongheFirstDay - 1065.0},
		{-1, 1, 1, JDOfGongheFirstDay - 731.0},
		{-1, 12, 30, JDOfGongheFirstDay - 367.0},
		{0, 1, 1, JDOfGongheFirstDay - 366.0},
		{0, 12, 31, JDOfGongheFirstDay - 1.0},
		{1, 1, 1, JDOfGongheFirstDay},
		{1, 1, 2, JDOfGongheFirstDay + 1.0},
		{1, 1, 30, JDOfGongheFirstDay + 29.0},
		{1, 2, 1, JDOfGongheFirstDay + 30.0},
		{1, 12, 30, JDOfGongheFirstDay + 364.0},
		{2, 2, 2, JDOfGongheFirstDay + 396.0},
		{4, 1, 1, JDOfGongheFirstDay + 1095.0},
		{4, 12, 31, JDOfGongheFirstDay + 1460.0},
		{5, 1, 1, JDOfGongheFirstDay + 1461.0},
		{100, 12, 30, JDOfGongheFirstDay + 36523.0},
		{101, 1, 1, JDOfGongheFirstDay + 36524.0},
		{500, 12, 31, JDOfGongheFirstDay + 182620.0},
		{501, 1, 1, JDOfGongheFirstDay + 182621.0},
		{2225, 11, 19, JDOfShuodanDongzhi},
	} {
		jd := GongheCalendarToJD(pair.y, pair.m, pair.d)
		assert.Equal(t, pair.jd, jd, "For date %04d-%02d-%02d expected %.1f got %.1f", pair.y, pair.m, pair.d, pair.jd, jd)
	}
}

func TestGCalToWCal(t *testing.T) {
	for _, pair := range []struct {
		gy, gm, gd int
		wy, wm, wd int
	}{
		{1, 1, 1, -841, 2, 12},
		{2225, 11, 19, 1384, 12, 13},
		{2752, 9, 7, 1911, 10, 10}, // not manual calculate yet
		{2819, 2, 2, 1978, 3, 4},   // not manual calculate yet
	} {
		y, m, d := GongheCalendarToWesternCalendar(pair.gy, pair.gm, pair.gd)
		assert.Equal(t, pair.wy, y, "For gcal date %04d-%02d-%02d expected year %d got %d", pair.gy, pair.gm, pair.gd, pair.wy, y)
		assert.Equal(t, pair.wm, m, "For gcal date %04d-%02d-%02d expected month %d got %d", pair.gy, pair.gm, pair.gd, pair.wm, m)
		assert.Equal(t, pair.wd, d, "For gcal date %04d-%02d-%02d expected day %d got %d", pair.gy, pair.gm, pair.gd, pair.wd, d)
	}
}

func TestWCalToGCal(t *testing.T) {
	for _, pair := range []struct {
		gy, gm, gd int
		wy, wm, wd int
	}{
		{1, 1, 1, -841, 2, 12},
		{2225, 11, 19, 1384, 12, 13},
		{2752, 9, 7, 1911, 10, 10}, // not manual calculate yet
		{2819, 2, 2, 1978, 3, 4},   // not manual calculate yet
	} {
		y, m, d := WesternCalendarToGongheCalendar(pair.wy, pair.wm, pair.wd)
		assert.Equal(t, pair.gy, y, "For wcal date %04d-%02d-%02d expected year %d got %d", pair.wy, pair.wm, pair.wd, pair.gy, y)
		assert.Equal(t, pair.gm, m, "For wcal date %04d-%02d-%02d expected month %d got %d", pair.wy, pair.wm, pair.wd, pair.gm, m)
		assert.Equal(t, pair.gd, d, "For wcal date %04d-%02d-%02d expected day %d got %d", pair.wy, pair.wm, pair.wd, pair.gd, d)
	}
}
