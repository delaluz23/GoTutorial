package main

import "fmt"

// Basic Data Types
func basicDataTypes() {
	/*
		Go has three basic data types:

		bool: represents a boolean value and is either true or false
		Numeric: represents integer types, floating point values, and complex types
		string: represents a string value

	*/
	var a bool = true     // Boolean
	var b int = 5         // Integer
	var c float32 = 3.14  // Floating point number 32 bit
	var e float64 = 3.123 // Floating point number 63 bit
	var d string = "Hi!"  // String

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float32:   ", c)
	fmt.Println("Float64: ", e)
	fmt.Println("String:  ", d)
}

// Boolean data type
func BooleanData() {
	//A boolean data type is declared with the bool keyword and can only take the values true or false.
	var b1 bool = true // typed declaration with initial value
	var b2 = true      // untyped declaration with initial value
	var b3 bool        // typed declaration without initial value
	b4 := true         // untyped declaration with initial value

	fmt.Println(b1) // Returns true
	fmt.Println(b2) // Returns true
	fmt.Println(b3) // Returns false
	fmt.Println(b4) // Returns true
}

// Integer data types

func IntegerData() {
	/*
		The integer data type has two categories:

		Signed integers - can store both positive and negative values
		Unsigned integers - can only store non-negative values
	*/
	//Signed integers, declared with one of the int keywords, can store both positive and negative values
	var x int = 500
	var y int = -4500
	fmt.Printf("Type: %T, value: %v", x, x)
	fmt.Printf("Type: %T, value: %v", y, y)
	/*
		Go has five keywords/types of signed integers:

		int	Depends on platform: 32 bits in 32 bit systems and 64 bit in 64 bit systems	-2147483648 to 2147483647 in 32 bit systems and-9223372036854775808 to 9223372036854775807 in 64 bit systems
		int8	8 bits/1 byte	-128 to 127
		int16	16 bits/2 byte	-32768 to 32767
		int32	32 bits/4 byte	-2147483648 to 2147483647
		int64	64 bits/8 byte	-9223372036854775808 to 9223372036854775807
	*/

	//Unsigned integers, declared with one of the uint keywords, can only store non-negative values
	var a uint = 500
	var b uint = 4500
	fmt.Printf("Type: %T, value: %v", a, a)
	fmt.Printf("Type: %T, value: %v", b, b)
	//if negative is given you get "cannot use *variable* (untyped int constant) as uint value in variable declaration (overflows)"
	/*
			Go has five keywords/types of unsigned integers

		uint	Depends on platform: 32 bits in 32 bit systems and 64 bit in 64 bit systems	0 to 4294967295 in 32 bit systems and 0 to 18446744073709551615 in 64 bit systems
		uint8	8 bits/1 byte	0 to 255
		uint16	16 bits/2 byte	0 to 65535
		uint32	32 bits/4 byte	0 to 4294967295
		uint64	64 bits/8 byte	0 to 18446744073709551615
	*/

}

// Float Data Type
func FloatData() {
	//The float data types are used to store positive and negative numbers with a decimal point
	/*
		The float data type has two keywords:

		float32	32 bits	-3.4e+38 to 3.4e+38.
		float64	64 bits	-1.7e+308 to +1.7e+308.
	*/

	//Float32
	var x float32 = 123.78
	var y float32 = 3.4e+38
	fmt.Printf("\nType: %T, value: %v\n", x, x)
	fmt.Printf("\nType: %T, value: %v", y, y)

	//Float64
	//The float64 data type can store a larger set of numbers than float32
	var z float64 = 1.7e+308
	fmt.Printf("\nType: %T, value: %v", z, z)
}

// String Data Type
func StringData() {
	//store a sequence of characters (text). String values must be surrounded by double quotes
	var txt1 string = "Hello!"
	var txt2 string
	txt3 := "World 1"

	fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
	fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
	fmt.Printf("Type: %T, value: %v\n", txt3, txt3)
}

//func main() {
//basicDataTypes()
//BooleanData()
//IntegerData()
//FloatData()
//StringData()
//}
