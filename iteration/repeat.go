package iteration

func Repeat(args string, repeatCount int) string {
	var result string

	for i := 0; i < repeatCount; i++ {
		result += args
	}

	return result
}
