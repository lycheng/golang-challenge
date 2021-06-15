package algorithm

import (
	"fmt"
	"strings"
)

// PrimeFactors from
//   https://www.codewars.com/kata/54d512e62a5e54c96200019e/train/go
func PrimeFactors(n int) string {
	rv := make([][]int, 0)
	f := 2
	for ; f*f <= n; f++ {
		cnt := 0
		for n%f == 0 {
			n /= f
			cnt++
		}

		if cnt > 0 {
			rv = append(rv, []int{f, cnt})
		}
	}

	if n > 1 {
		rv = append(rv, []int{n, 1})
	}

	sb := strings.Builder{}
	for _, item := range rv {
		f := item[0]
		cnt := item[1]
		if cnt > 1 {
			sb.WriteString(fmt.Sprintf("(%d**%d)", f, cnt))
		} else {
			sb.WriteString(fmt.Sprintf("(%d)", f))
		}
	}
	return sb.String()
}
