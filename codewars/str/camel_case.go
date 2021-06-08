package str

import "strings"

// CamelCase from
//   https://www.codewars.com/kata/587731fda577b3d1b0001196/train/go
func CamelCase(s string) string {
	return strings.Join(stringMap(strings.Split(s, " "), strings.Title), "")
}

func stringMap(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
