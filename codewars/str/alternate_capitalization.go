package str

import "bytes"

// Capitalize from
//   https://www.codewars.com/kata/59cfc000aeb2844d16000075/train/go
func Capitalize(st string) []string {
	rv := []bytes.Buffer{{}, {}}
	for i, c := range st {
		uc := c - 'a' + 'A'
		rv[i%2].WriteRune(uc)
		rv[(i+1)%2].WriteRune(c)
	}
	return []string{rv[0].String(), rv[1].String()}
}
