package str

// FirstNonRepeating from
//   https://www.codewars.com/kata/52bc74d4ac05d0945d00054e/train/go
func FirstNonRepeating(str string) string {
	m := make(map[rune]int)
	for _, c := range str {
		m[lowerCase(c)]++
	}

	for _, c := range str {
		if m[lowerCase(c)] == 1 {
			return string(c)
		}
	}
	return ""
}

func lowerCase(c rune) rune {
	if c >= 'A' && c <= 'Z' {
		c = c - 'A' + 'a'
	}
	return c
}
