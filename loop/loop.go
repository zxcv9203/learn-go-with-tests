package loop

const repeatCount = 5

func Repeat(str string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += str
	}
	return repeated
}
