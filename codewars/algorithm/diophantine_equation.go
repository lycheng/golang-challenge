package algorithm

import "math"

// Solequa from
//   https://www.codewars.com/kata/554f76dca89983cc400000bb/train/go
func Solequa(n int) [][]int {
	rv := make([][]int, 0)

	limit := int(math.Sqrt(float64(n))) + 1
	for i := 1; i < limit; i++ {
		if n%i != 0 {
			continue
		}
		j := n / i

		if j < i {
			break
		}

		if (i+j)%2 != 0 || (j-i)%4 != 0 {
			continue
		}
		rv = append(rv, []int{(i + j) / 2, (j - i) / 4})
	}
	return rv
}
