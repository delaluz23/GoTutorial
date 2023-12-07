package main

import "fmt"

func Go_Array() {
	/*
		SYNTAX :
		var array_name = [length]datatype{values} // here length is defined
		or
		var array_name = [...]datatype{values} // here length is inferred
		or
		array_name :=
	*/

	//defined length example
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)

	//inferred length example
	var arr11 = [...]int{1, 2, 3}
	arr22 := [...]int{4, 5, 6, 7, 8}

	fmt.Println(arr11)
	fmt.Println(arr22)

	//how to change index value in array

	prices := [3]int{10, 20, 30}

	prices[2] = 50
	fmt.Println("prices: ", prices)

	//Array initialization
	//If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type
	Arr1 := [5]int{}              //not initialized
	Arr2 := [5]int{1, 2}          //partially initialized
	Arr3 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println("NOT INITIALIZED: ", Arr1)
	fmt.Println("PARTIALLY INITIALIZED: ", Arr2)
	fmt.Println("FULLY INITIALIZED: ", Arr3)

	//initializing specific elements
	array1 := [5]int{1: 10, 2: 40}

	fmt.Println("initialized index1 and index2 : ", array1)

	//len() function is used to find the length of an array
	//defined array
	cars := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	//inferred array
	nums := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println("length of cars array: ", len(cars))
	fmt.Println("length of nums array: ", len(nums))
}

//func main() {
//	Go_Array()
//}
