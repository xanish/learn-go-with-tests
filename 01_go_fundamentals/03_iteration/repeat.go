package _3_iteration

const defaultTimes = 5

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
