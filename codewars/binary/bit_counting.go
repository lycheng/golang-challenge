package binary

// CountBits from
//   https://www.codewars.com/kata/526571aae218b8ee490006f4/train/go
func CountBits(n uint) int {
	cnt := 0
	for ; n > 0; n = n >> 1 {
		cnt += int(n & 1)
	}
	return cnt
}
