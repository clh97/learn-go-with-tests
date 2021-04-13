package iteration

func Repeat(toRepeat string, qty int) (repeated string) {
	for i := 0; i < qty; i++ {
		repeated += toRepeat
	}
	return repeated
}
