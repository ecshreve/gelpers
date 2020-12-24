package gelpers_test

import (
	"fmt"

	"github.com/ecshreve/gelpers"
)

func ExampleStringPtr() {
	sPtr := gelpers.StringPtr("hello")
	fmt.Printf("%T", sPtr)
	// Output: *string
}

func ExampleStringVal() {
	s := "hello"
	sPtr := &s
	sVal := gelpers.StringVal(sPtr)
	fmt.Printf("%T", sVal)
	// Output: string
}

func ExampleStringVal_second() {
	var sPtr *string
	sVal := gelpers.StringVal(sPtr)
	fmt.Println(sVal == "")
	// Output: true
}
