package algorithm

// Tickets from
//   https://www.codewars.com/kata/555615a77ebc7c2c8a0000b8/train/go
func Tickets(peopleInLine []int) string {
	cnt25 := 0
	cnt50 := 0

	for _, pay := range peopleInLine {
		switch pay {
		case 25:
			cnt25++
		case 50:
			if cnt25 < 1 {
				return "NO"
			}
			cnt25--
			cnt50++
		default:
			if cnt25 >= 1 && cnt50 >= 1 {
				cnt25--
				cnt50--
			} else if cnt25 >= 3 {
				cnt25 -= 3
			} else {
				return "NO"
			}
		}
	}
	return "YES"
}
