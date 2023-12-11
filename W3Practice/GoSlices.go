package main

import "fmt"

/*
	Slices are similar to arrays, but are more powerful and flexible.

	Like arrays, slices are also used to store multiple values of the same type in a single variable.

	However, unlike arrays, the length of a slice can grow and shrink as you see fit.

	In Go, there are several ways to create a slice:

	Using the []datatype{values} format
	Create a slice from an array
	Using the make() function
*/

// Create a Slice With []datatype{values}
func SliceWDatatype() {
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
}

// Create a Slice with a array
func SliceWArray() {
	/*
		SYNTAX DIFFERENCE BETWEEN SLICE W/ ARRAY VS ARRAY
		var myarray = [length]datatype{values} // An array
		myslice := myarray[start:end] // A slice made from the array
	*/
	//first create array
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	//second pust start and end
	myslice := arr1[0:len(arr1)]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))
}

//Create a Slice With The make() Function

func SliceWMakeFunc() {
	/*
		SYNTAX
		slice_name := make([]type, length, capacity)

		If the capacity parameter is not defined, it will be equal to length.
	*/
	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	// with omitted capacity
	myslice2 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))
}

// MODIFY SLICES
func modSlices() {
	prices := []int{10, 20, 30}
	prices[2] = 50
	fmt.Println(prices[0])
	fmt.Println(prices[2])

	//APPENDING TO A SLICE
	myslice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	//APPENDING ONE SLICE TO ANOTHER

}
func main() {
	//SliceWDatatype()
	//SliceWArray()
	//SliceWMakeFunc()
	modSlices()
}
