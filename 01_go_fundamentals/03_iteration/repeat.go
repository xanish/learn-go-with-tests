package _3_iteration

const defaultTimes = 5

// Repeat takes a character string and repeats it 'times' number of times.
// If 'times' is 0, it defaults to 'defaultTimes'.
func Repeat(character string, times int) string {
	if times == 0 {
		times = defaultTimes
	}

	var repeated string
	for i := 0; i < times; i++ {
		repeated = repeated + character
	}
	return repeated
}
