package algorithm

// Arithmetic from
//   https://www.codewars.com/kata/583f158ea20cfcbeb400000a/train/go
func Arithmetic(a int, b int, operator string) int {
	m := map[string]int{
		"add":      a + b,
		"subtract": a - b,
		"multiply": a * b,
		"divide":   a / b,
	}
	return m[operator]
}
