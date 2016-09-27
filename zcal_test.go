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
