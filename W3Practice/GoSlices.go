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

}

// APPENDING ONE SLICE TO ANOTHER
// Append all the elements of one slice to another slice, use the append()function
func appendingSlice() {
	myslice1 := []int{1, 2, 3}
	myslice2 := []int{4, 5, 6}
	myslice3 := append(myslice1, myslice2...)

	fmt.Printf("myslice3=%v\n", myslice3)
	fmt.Printf("length=%d\n", len(myslice3))
	fmt.Printf("capacity=%d\n", cap(myslice3))
}

// CHANGING LENGTH OF SLICE
func sliceLengthChange() {
	arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
	myslice1 := arr1[1:5]                 // Slice array
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = arr1[1:3] // Change length by re-slicing the array
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
}

//MEMORY EFFICIENCY
/*
When using slices, Go loads all the underlying elements into the memory.
If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.
The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program.
*/
func copySlice() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}

//func main() {
//SliceWDatatype()
//SliceWArray()
//SliceWMakeFunc()
//modSlices()
//appendingSlice()
//sliceLengthChange()
//copySlice()
//}
