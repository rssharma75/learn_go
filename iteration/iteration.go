package iteration

func Repeat(a string, count int) string {
	str := ""

	for i := 0; i < count; i++ {
		str += a
	}
	return str
}
