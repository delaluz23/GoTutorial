package main

import (
	"fmt"
)

func constVariables() {

	//TWO TYPES OF CONSTANTS
	/*
		Typed constants are declared with a defined type: const A int = 1
		Untyped constants are declared without a type: const A = 1
	*/
	const NUMBER_UNTYPED = 1
	fmt.Println(NUMBER_UNTYPED)
	const NUMBER_TYPED int = 2
	fmt.Println(NUMBER_TYPED)

	//Multiple constant declaration
	const (
		A int = 1
		B     = 3.14
		C     = "Hi!"
	)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}

// Constants value cannot be changed
func main() {
	constVariables()
}

//CONSTANT RULES
/*
	Constant names follow the same naming rules as variables
	Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
	Constants can be declared both inside and outside of a function
*/
