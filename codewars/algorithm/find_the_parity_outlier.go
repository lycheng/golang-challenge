package algorithm

// FindOutlier from
//   https://www.codewars.com/kata/5526fc09a1bbd946250002dc/train/go
func FindOutlier(integers []int) int {

	m := make(map[int][]int)
	m[0] = make([]int, 0)
	m[1] = make([]int, 0)

	for _, n := range integers {
		isOdd := n % 2
		if isOdd < 0 {
			isOdd *= -1
		}
		m[isOdd] = append(m[isOdd], n)
		nn, gotIt := getTheOutlier(m)
		if gotIt {
			return nn
		}
	}
	return -1
}

func getTheOutlier(m map[int][]int) (int, bool) {
	if len(m[0]) == 1 && len(m[1]) > 1 {
		return m[0][0], true
	}

	if len(m[0]) > 1 && len(m[1]) == 1 {
		return m[1][0], true
	}
	return -1, false
}
