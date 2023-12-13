package main

import "fmt"

//func main() {
//	IfStatment()
//	IfElseStatment()
//	ElseIfStatment()
//	NestedIf()
//}

func NestedIf() {
	num := 20
	if num >= 10 {
		fmt.Println("Num is more than 10.")
		if num > 15 {
			fmt.Println("Num is also more than 15.")
		}
	} else {
		fmt.Println("Num is less than 10.")
	}
}

func ElseIfStatment() {
	time := 22
	if time < 10 {
		fmt.Println("Good morning.")
	} else if time < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}
}

func IfElseStatment() {
	time := 20
	if time < 18 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}
}

func IfStatment() {
	if 20 > 18 {
		fmt.Println("20 is greater than 18")
	}
}

func Conditions() {

}
