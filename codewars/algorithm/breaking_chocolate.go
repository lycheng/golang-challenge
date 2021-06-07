package algorithm

// BreakChocolate from
//   https://www.codewars.com/kata/534ea96ebb17181947000ada/train/go
func BreakChocolate(n, m int) int {
	if n*m <= 1 {
		return 0
	}
	return n*m - 1
}
