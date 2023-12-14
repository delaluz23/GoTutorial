package main

import "fmt"

/*
Maps are used to store data values in key:value pairs.
Each element in a map is a key:value pair.
A map is an unordered and changeable collection that does not allow duplicates.
The length of a map is the number of its elements. You can find it using the len() function.
The default value of a map is nil.
Maps hold references to an underlying hash table.
Go has multiple ways for creating maps.
*/
//func main() {
/*
	valid key types are:
		Booleans
		Numbers
		Strings
		Arrays
		Pointers
		Structs
		Interfaces (as long as the dynamic type supports equality)
	Invalid key types are:
		Slices
		Maps
		Functions
*/
//varMap()
//makeMap()
//emptyMap()
//accessing map elements
//accessMap()
//}

func varMap() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
}

func makeMap() {
	var a = make(map[string]string) // The map is empty now
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
	// a is no longer empty
	b := make(map[string]int)
	b["Oslo"] = 1
	b["Bergen"] = 2
	b["Trondheim"] = 3
	b["Stavanger"] = 4

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
}

func emptyMap() {
	var a = make(map[string]string)
	var b map[string]string

	fmt.Println(a == nil)
	fmt.Println(b == nil)
}

func accessMap() {
	var a = make(map[string]string)
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"

	fmt.Printf(a["brand"])
}
