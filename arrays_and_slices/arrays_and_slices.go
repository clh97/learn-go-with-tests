package arrays_and_slices

func Sum(numbers [5]int) (result int) {
	for _, num := range numbers {
		result += num
	}
	return result
}
