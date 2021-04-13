package arrays_and_slices

func Sum(numbers []int) (result int) {
	for _, num := range numbers {
		result += num
	}
	return result
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
