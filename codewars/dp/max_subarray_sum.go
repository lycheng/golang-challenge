package dp

// MaximumSubarraySum from
//   https://www.codewars.com/kata/54521e9ec8e60bc4de000d6c/train/go
func MaximumSubarraySum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	max := 0
	sum := 0

	for _, n := range numbers {
		sum += n
		if sum < 0 {
			sum = 0
		} else if sum > max {
			max = sum
		}
	}
	return max
}
