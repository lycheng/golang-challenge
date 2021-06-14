package algorithm

// CountKprimes from
//   from https://www.codewars.com/kata/5726f813c8dcebf5ed000a6b/train/go
func CountKprimes(k, start, nd int) []int {
	rv := make([]int, 0)
	for i := start; i <= nd; i++ {
		if kPrime(i) == k {
			rv = append(rv, i)
		}
	}
	if len(rv) == 0 {
		return nil
	}
	return rv
}

func kPrime(n int) (k int) {
	i := 2
	for i*i <= n {
		for n%i == 0 {
			n /= i
			k++
		}
		i++
	}
	if n > 1 {
		k++
	}
	return
}

// Puzzle from
//   from https://www.codewars.com/kata/5726f813c8dcebf5ed000a6b/train/go
func Puzzle(s int) int {
	rv := 0

	a := CountKprimes(1, 0, s)
	b := CountKprimes(3, 0, s)
	c := CountKprimes(7, 0, s)

	for _, i := range a {
		for _, j := range b {
			for _, k := range c {
				if i+j+k == s {
					rv++
				}
			}
		}
	}
	return rv
}
