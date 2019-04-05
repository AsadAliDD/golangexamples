package golangexamples

//Slice concataion
func ConcatSlice(sliceToConcat []byte) string {
	var dummy string
	for index := 0; index < len(sliceToConcat)-1; index++ {
		dummy = dummy + string(sliceToConcat[index]) + "-"
	}
	dummy = dummy + string(sliceToConcat[len(sliceToConcat)-1])
	return dummy
}

func Encrypt(sliceToConcat []byte, ceaserCount int) string {

	// var dummy strings
	return "HELLO"
}
