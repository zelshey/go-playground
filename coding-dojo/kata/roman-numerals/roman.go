package roman

import "strings"

type RomanNumeral struct {
	Arabic int
	Roman  string
}

func getCharacters() []RomanNumeral {
	return []RomanNumeral{
		{Arabic: 1000, Roman: "M"},
		{Arabic: 500, Roman: "D"},
		{Arabic: 100, Roman: "C"},
		{Arabic: 50, Roman: "L"},
		{Arabic: 10, Roman: "X"},
		{Arabic: 5, Roman: "V"},
		{Arabic: 1, Roman: "I"},
	}
}

func ArabicToRoman(digit int) string {
	characters := getCharacters()
	ret := ""

	for i, c := range characters {
		for digit >= c.Arabic {
			ret += c.Roman
			digit -= c.Arabic
		}

		for _, n := range characters[i+1:] {
			value := c.Arabic - n.Arabic
			for digit >= value && value != n.Arabic {
				ret += n.Roman + c.Roman
				digit -= value
			}
		}

	}

	return ret
}

func RomanToArabic(roman string) int {
	ret := 0
	characters := getCharacters()

	for i, c := range characters {
		for _, p := range characters[:i] {
			pre := c.Roman + p.Roman
			for strings.HasPrefix(roman, pre) {
				roman = strings.TrimPrefix(roman, pre)
				ret += p.Arabic - c.Arabic
			}
		}

		for strings.HasPrefix(roman, c.Roman) {
			roman = strings.TrimPrefix(roman, c.Roman)
			ret += c.Arabic
		}
	}

	return ret
}
