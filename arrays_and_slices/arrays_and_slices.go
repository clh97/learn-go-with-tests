package arrays_and_slices

func Sum(numbers [5]int) (result int) {
	for i := 0; i < 5; i++ {
		result += numbers[i]
	}
	return result
}