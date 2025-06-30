package roman

import "log"

func ArabicToRoman(digit int) string {
	characters := []struct {
		Arabic int
		Roman  string
	}{
		{Arabic: 1000, Roman: "M"},
		{Arabic: 500, Roman: "D"},
		{Arabic: 100, Roman: "C"},
		{Arabic: 50, Roman: "L"},
		{Arabic: 10, Roman: "X"},
		{Arabic: 5, Roman: "V"},
		{Arabic: 1, Roman: "I"},
	}

	ret := ""

	for {
		if digit < 0 {
			log.Fatal("Invalid value during calculation")
		}

		if digit == 0 {
			break
		}

		for _, c := range characters {
			if digit >= c.Arabic {
				ret += c.Roman
				digit -= c.Arabic
				break
			}
		}
	}

	return ret
}
