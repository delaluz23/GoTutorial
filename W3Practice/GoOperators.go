package main

import "fmt"

// ARITHMETIC OPERATORS
func arithmeticOperators() {
	/*
		+	Addition	Adds together two values	x + y
		-	Subtraction	Subtracts one value from another	x - y
		*	Multiplication	Multiplies two values	x * y
		/	Division	Divides one value by another	x / y
		%	Modulus	Returns the division remainder	x % y
		++	Increment	Increases the value of a variable by 1	x++
		--	Decrement	Decreases the value of a variable by 1	x--
	*/
	var x = 10
	for i := 0; i <= x; i++ {
		fmt.Println(i)
	}
	for i := 20; i >= x; i-- {
		fmt.Println(i)
	}
}

// ASSIGNMENT OPERATORS
func assignmentOperators() {
	/*
		=	x = 5	x = 5
		+=	x += 3	x = x + 3
		-=	x -= 3	x = x - 3
		*=	x *= 3	x = x * 3
		/=	x /= 3	x = x / 3
		%=	x %= 3	x = x % 3
		&=	x &= 3	x = x & 3
		|=	x |= 3	x = x | 3
		^=	x ^= 3	x = x ^ 3
		>>=	x >>= 3	x = x >> 3
		<<=	x <<= 3	x = x << 3
	*/
	var x = 5
	fmt.Println(x, "= x and ^ 5 =")
	x ^= 5
	fmt.Println(x)

}

// COMPARISON OPERATORS
func comparisonOperators() {
	/*
		==	Equal to	x == y
		!=	Not equal	x != y
		>	Greater than	x > y
		<	Less than	x < y
		>=	Greater than or equal to	x >= y
		<=	Less than or equal to	x <= y
	*/
	var x = 5
	var y = 3
	fmt.Println(x > y) // returns 1 (true) because 5 is greater than 3
}

// LOGICAL OPERATORS
func logicalOperators() {
	/*
		&& 	Logical and	Returns true if both statements are true	x < 5 &&  x < 10
		|| 	Logical or	Returns true if one of the statements is true	x < 5 || x < 4
		!	Logical not	Reverse the result, returns false if the result is true	!(x < 5 && x < 10)
	*/
}

// BITWISE OPERATORS
func bitwiseOperators() {
	/*
		& 	AND	Sets each bit to 1 if both bits are 1	x & y
		|	OR	Sets each bit to 1 if one of two bits is 1	x | y
		 ^	XOR	Sets each bit to 1 if only one of two bits is 1	x ^ b
		<<	Zero fill left shift	Shift left by pushing zeros in from the right	x << 2
		>>	Signed right shift	Shift right by pushing copies of the leftmost bit in from the left, and let the rightmost bits fall off	x >> 2
	*/
}
func main() {
	//arithmeticOperators()
	//assignmentOperators()
	//comparisonOperators()
	//logicalOperators()
	//bitwiseOperators()
}
