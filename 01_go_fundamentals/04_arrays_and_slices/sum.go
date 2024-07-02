package _4_arrays_and_slices

func Sum(numbers []int) int {
	var result int
	for _, number := range numbers {
		result += number
	}
	return result
}

func SumAll(numbers ...[]int) []int {
	var result []int

	for _, slice := range numbers {
		result = append(result, Sum(slice))
	}

	return result
}
