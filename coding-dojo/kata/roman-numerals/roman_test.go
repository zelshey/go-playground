package roman

import (
	"fmt"
	"testing"
)

func TestArabicToRoman(t *testing.T) {
	cases := []struct {
		Arabic int
		Roman  string
	}{
		{Arabic: 1, Roman: "I"},
		{Arabic: 2, Roman: "II"},
		{Arabic: 3, Roman: "III"},
		{Arabic: 4, Roman: "IV"},
		{Arabic: 5, Roman: "V"},
		{Arabic: 6, Roman: "VI"},
		{Arabic: 7, Roman: "VII"},
		{Arabic: 8, Roman: "VIII"},
		// {Arabic: 9, Roman: "IX"},
		{Arabic: 10, Roman: "X"},
		// {Arabic: 14, Roman: "XIV"},
		{Arabic: 18, Roman: "XVIII"},
		{Arabic: 20, Roman: "XX"},
		// {Arabic: 39, Roman: "XXXIX"},
		// {Arabic: 40, Roman: "XL"},
		// {Arabic: 47, Roman: "XLVII"},
		// {Arabic: 49, Roman: "XLIX"},
		{Arabic: 50, Roman: "L"},
		{Arabic: 100, Roman: "C"},
		// {Arabic: 90, Roman: "XC"},
		// {Arabic: 400, Roman: "CD"},
		{Arabic: 500, Roman: "D"},
		// {Arabic: 900, Roman: "CM"},
		{Arabic: 1000, Roman: "M"},
		// {Arabic: 1984, Roman: "MCMLXXXIV"},
		// {Arabic: 3999, Roman: "MMMCMXCIX"},
		// {Arabic: 2014, Roman: "MMXIV"},
		{Arabic: 1006, Roman: "MVI"},
		// {Arabic: 798, Roman: "DCCXCVIII"},
	}

	for _, tc := range cases {
		testName := fmt.Sprintf("Convert %d to %s", tc.Arabic, tc.Roman)
		t.Run(testName, func(t *testing.T) {
			numeral := ArabicToRoman(tc.Arabic)
			if numeral != tc.Roman {
				t.Errorf("got %s, want %s", numeral, tc.Roman)
			}
		})
	}
}
