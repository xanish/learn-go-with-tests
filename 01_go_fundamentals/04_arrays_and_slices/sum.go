package _4_arrays_and_slices

// Sum calculates the sum of all numbers in the given slice of integers.
func Sum(numbers []int) int {
	var result int
	for _, number := range numbers {
		result += number
	}
	return result
}

// SumAll calculates the sum of each slice of integers provided as variadic arguments.
// It returns a slice containing the sums of each input slice.
func SumAll(numbers ...[]int) []int {
	var result []int

	for _, slice := range numbers {
		result = append(result, Sum(slice))
	}

	return result
}

// SumAllTails calculates the sum of all tails (sub-slices excluding the first element)
// of each slice of integers provided as variadic arguments.
// It returns a slice containing the sums of tails for each input slice.
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
