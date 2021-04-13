package arrays_and_slices

func Sum(numbers []int) (result int) {
	for _, num := range numbers {
		result += num
	}
	return result
}

func SumAll(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
