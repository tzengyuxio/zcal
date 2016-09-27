package zcal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/tzengyuxio/zcal"
)

func TestGregorianYearToStemBranch(t *testing.T) {
	var tests = []struct {
		year       int
		stemBranch string
	}{
		{-2697, "甲子"}, // 黃帝紀元元年
		{-841, "庚申"},  // 共和元年，中國確切紀年之始
		{-221, "庚辰"},  // 秦始皇二十六年，統一中國
		{4, "甲子"},     // 黃帝紀元45甲子(2700年)之後
		{1384, "甲子"},  // 洪武十七年
		{1894, "甲午"},  // 甲午戰爭
		{1912, "壬子"},  // 中華民國元年
		{2012, "壬辰"},
	}

	for _, pair := range tests {
		stemBranch := GregorianYearToStemBranch(pair.year)
		assert.Equal(t, pair.stemBranch, stemBranch, "For year %d expected %s got %s", pair.year, pair.stemBranch, stemBranch)
	}
}
