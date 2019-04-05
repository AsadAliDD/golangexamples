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
func Encrypt(sliceToConcat []byte, ceaserCount int) string {
	// var dummy strings
	var dummy string
	for index := 0; index < len(sliceToConcat); index++ {
		dummy = dummy + string(sliceToConcat[index]+byte(ceaserCount))
	}
	return dummy
}

//EZGreetings
func EZGreetings(dummy string) string {
	return greetings.PrintGreetings(dummy)
}
