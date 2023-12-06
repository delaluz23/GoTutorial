package main

import "fmt"

// Three ways to output
func Output() {
	var i, j = "Hello", "World"

	//1. PRINT
	fmt.Print(i, j)
	//if you wanted a space between i and j use: i," ",j

	//If we want to print the arguments in new lines, we need to use \n

	fmt.Print("\n", i, "\n")
	fmt.Print(j, "\n")

	//Print() inserts a space between the arguments if neither are strings:
	var a, b = 1, 2
	fmt.Print(a, b, "\n")

	//2. PRINTLN

	//whitespace is added between the arguments, and a newline is added at the end
	fmt.Println(i, j)

	//3. PRINTF
	/*
		%v	Prints the value in the default format
		%#v	Prints the value in Go-syntax format
		%T	Prints the type of the value
		%%	Prints the % sign
	*/
	var (
		c string  = "JOSE"
		d float64 = 10.1
		e string  = "DOG"
	)
	fmt.Printf("c has the value: %v and type type: %T\\n", c, c)
	fmt.Printf("d has the value: %v and type type: %T\\n", d, d, "\n")
	fmt.Printf("\n%%v of e = %v", e)
	fmt.Printf("\n%%#v of e = %#v", e)
	fmt.Printf("\n%%T of e = %T", e)

	//integer formating verbs
	/*
		%b	Base 2
		%d	Base 10
		%+d	Base 10 and always show sign
		%o	Base 8
		%O	Base 8, with leading 0o
		%x	Base 16, lowercase
		%X	Base 16, uppercase
		%#x	Base 16, with leading 0x
		%4d	Pad with spaces (width 4, right justified)
		%-4d	Pad with spaces (width 4, left justified)
		%04d	Pad with zeroes (width 4
	*/
	var number int = 2
	fmt.Println("\nFormating with the number 2")
	fmt.Printf("\n%b\n", number)
	fmt.Printf("%d\n", number)
	fmt.Printf("%+d\n", number)
	fmt.Printf("%o\n", number)
	fmt.Printf("%O\n", number)
	fmt.Printf("%x\n", number)
	fmt.Printf("%X\n", number)
	fmt.Printf("%#x\n", number)
	fmt.Printf("%4d\n", number)
	fmt.Printf("%-4d\n", number)
	fmt.Printf("%04d\n", number)

	//String formatting verbs
	/*
		%s	Prints the value as plain string
		%q	Prints the value as a double-quoted string
		%8s	Prints the value as plain string (width 8, right justified)
		%-8s	Prints the value as plain string (width 8, left justified)
		%x	Prints the value as hex dump of byte values
		% x	Prints the value as hex dump with spaces
	*/
	var txt = "Hello"

	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

	//Boolean formatting verbs
	/*
		%t	Value of the boolean operator in true or false format (same as using %v)
	*/
	var x = true
	var z = false

	fmt.Printf("%t\n", x)
	fmt.Printf("%t\n", z)

	//Float formatting verb
	/*
		%e	Scientific notation with 'e' as exponent
		%f	Decimal point, no exponent
		%.2f	Default width, precision 2
		%6.2f	Width 6, precision 2
		%g	Exponent as needed, only necessary digits
	*/
	var g = 3.141

	fmt.Printf("%e\n", g)
	fmt.Printf("%f\n", g)
	fmt.Printf("%.2f\n", g)
	fmt.Printf("%6.2f\n", g)
	fmt.Printf("%g\n", g)
}

//func main() {
//	Output()
//}
