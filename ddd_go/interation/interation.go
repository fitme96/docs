package interation

func Repeat(character string, conut int) string {
	var repeated string

	for i := 0; i < conut; i++ {
		repeated += character
	}

	return repeated
}
