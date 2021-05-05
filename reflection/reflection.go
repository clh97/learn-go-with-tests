package reflection

func walk(x interface{}, fn func(string)) {
	fn("the quick brown fox jumps over the lazy dog")
}
