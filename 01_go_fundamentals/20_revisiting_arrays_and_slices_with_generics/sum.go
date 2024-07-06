package _0_revisiting_arrays_and_slices_with_generics

// Sum calculates the sum of all numbers in the given slice of integers.
func Sum(numbers []int) int {
	adder := func(accumulator, x int) int { return accumulator + x }
	return Reduce(numbers, adder, 0)
}

// SumAll calculates the sum of each slice of integers provided as variadic arguments.
// It returns a slice containing the sums of each input slice.
func SumAll(numbers ...[]int) []int {
	sum := func(accumulator, slice []int) []int {
		if len(slice) == 0 {
			return append(accumulator, 0)
		} else {
			return append(accumulator, Sum(slice))
		}
	}

	return Reduce(numbers, sum, []int{})
}

// SumAllTails calculates the sum of all tails (sub-slices excluding the first element)
// of each slice of integers provided as variadic arguments.
// It returns a slice containing the sums of tails for each input slice.
func SumAllTails(numbers ...[]int) []int {
	sumTail := func(accumulator, slice []int) []int {
		if len(slice) == 0 {
			return append(accumulator, 0)
		} else {
			tail := slice[1:]
			return append(accumulator, Sum(tail))
		}
	}

	return Reduce(numbers, sumTail, []int{})
}
