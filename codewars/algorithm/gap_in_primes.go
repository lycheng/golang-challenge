package algorithm

// Gap from
//   https://www.codewars.com/kata/561e9c843a2ef5a40c0000a4/train/go
func Gap(g, m, n int) []int {
	l := n - m + 1
	flag := make([]bool, l)
	for i := 0; i < l; i++ {
		flag[i] = isPrime(i + m)
	}

	for i := 0; i < l; i++ {
		if !flag[i] {
			continue
		}
		for j := i + 1; j < l; j++ {
			if flag[j] {
				if j-i == g {
					return []int{i + m, j + m}
				}
				break
			}
		}
	}
	return nil
}

// isPrime from
//   https://en.wikipedia.org/wiki/Primality_test#Python_code
func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i < n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
