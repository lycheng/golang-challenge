package str

import "strings"

// https://www.codewars.com/kata/54bf1c2cd5b56cc47f0007a1/train/go
func duplicateCount(s1 string) int {
	s := strings.ToLower(s1)
	m := make(map[rune]int)
	rv := 0

	for _, c := range s {
		if m[c]++; m[c] == 2 {
			rv++
		}
	}
	return rv
}
