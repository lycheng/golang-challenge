package algorithm

import (
	"strings"
)

const maxFacToCal = 20

var factMap = map[int]int{0: 0, 1: 1}

func init() {
	for i := 2; i <= maxFacToCal; i++ {
		factMap[i] = i * factMap[i-1]
	}
}

func getBase(i int) rune {
	switch {
	case i >= 0 && i <= 9:
		return rune(i + '0')
	case i > 9:
		return rune(int('A') + i - 10)
	}
	return '0'
}

func baseToNum(r rune) int {
	switch {
	case r >= '0' && r <= '9':
		return int(r - '0')
	case r >= 'A':
		return int(r-'A') + 10
	}
	return 0
}

// Dec2FactString from
//   https://www.codewars.com/kata/54e320dcebe1e583250008fd/train/go
func Dec2FactString(nb int) string {
	buff := strings.Builder{}
	for i := maxFacToCal; i > 0; i-- {
		if nb < factMap[i] {
			if buff.Len() > 0 {
				buff.WriteRune('0')
			}
			continue
		}
		j := nb / factMap[i]
		nb -= factMap[i] * j
		buff.WriteRune(getBase(j))
	}
	buff.WriteRune('0')
	return buff.String()
}

// FactString2Dec from
//   https://www.codewars.com/kata/54e320dcebe1e583250008fd/train/go
func FactString2Dec(str string) int {
	runes := []rune(str)
	rv := 0
	l := len(runes)
	for i := l - 1; i >= 0; i-- {
		j := l - i - 1
		n := baseToNum(runes[i])
		rv += n * factMap[j]
	}
	return rv
}
