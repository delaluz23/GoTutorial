package main

import "fmt"

//func main() {
//	familyName("Liam")
//	familyName2("Liam", 3)
//
//	//returns func
//	//If we (for some reason) do not want to use some of the returned values, we can add an underscore (_), to omit this value.
//	_, b := myFunction(5, "Hello")
//	fmt.Println(b)
//
//	//recursive func
//	testcount(1)
//}

// CREATING FUNCTION USE func

// Parameters and arguments
func familyName(fname string) {
	fmt.Println("Hello", fname, "Refsnes")
}

// Multiple parameters
func familyName2(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

// Multiple returns
func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

// Recursion functions
func testcount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
}
