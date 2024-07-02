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

func SumAllTails(numbers ...[]int) []int {
	var result []int
	for _, slice := range numbers {
		if len(slice) == 0 {
			result = append(result, 0)
		} else {
			tail := slice[1:]
			result = append(result, Sum(tail))
		}
	}

	return result
}
