package golangexamples

import (
	"github.com/ehteshamz/greetings"
)

//ConcatSlice Returns a Concatenanted byte array into string
func ConcatSlice(sliceToConcat []byte) string {
	var dummy string
	for index := 0; index < len(sliceToConcat)-1; index++ {
		dummy = dummy + string(sliceToConcat[index]) + "-"
	}
	dummy = dummy + string(sliceToConcat[len(sliceToConcat)-1])
	return dummy
}

//Encrypt
func Encrypt(sliceToConcat []byte, ceaserCount int) {

	for index := 0; index < len(sliceToConcat); index++ {
		sliceToConcat[index] = sliceToConcat[index] + byte(ceaserCount)
	}
}

//EZGreetings
func EZGreetings(dummy string) string {
	return greetings.PrintGreetings(dummy)
}
