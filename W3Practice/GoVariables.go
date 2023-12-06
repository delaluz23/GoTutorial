package main

import (
	"fmt"
	_ "fmt"
)

// two ways to declare variables

func variables() {
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred

	// Can only be used inside of functions
	x := 2 //type is inferred

	//Below shows the data type and value
	fmt.Println("Value of student1:", student1)
	fmt.Printf("Type of student1: %T\n", student1)
	fmt.Println("-------")
	fmt.Println("Value of student2:", student2)
	fmt.Printf("Type of student2: %T\n", student2)
	fmt.Println("-------")
	fmt.Println("Value of x:", x)
	fmt.Printf("Type of x: %T\n", x)
}

//FUNCTION FOR MULTIPLE VARIABLES

func multVariables() {
	////can declare multiple variables on the same line
	////If you use the type keyword, it is only possible to declare one type of variable per line.
	//var a, b, c, d int = 1, 3, 5, 7
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	//fmt.Println(d)

	//If the type keyword is not specified, you can declare different types of variables in the same line:
	var a, b = 6, "Hello"
	c, d := 7, "World!"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}

// FUNCTION FOR BLOCK VARIABLES

func blockVariables() {
	var (
		a int
		b int    = 1
		c string = "hello"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

//func main() {
//	//variables()
//	//multVariables()
//	blockVariables()
//}

//VARIABLE NAMING RULES
/*
	A variable name must start with a letter or an underscore character (_)
	A variable name cannot start with a digit
	A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
	Variable names are case-sensitive (age, Age and AGE are three different variables)
	There is no limit on the length of the variable name
	A variable name cannot contain spaces
	The variable name cannot be any Go keywords
*/
