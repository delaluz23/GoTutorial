package main

import (
	"fmt"
	_ "fmt"
)

// two ways to declare variables

func variables() {
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	fmt.Printf(student1)
	fmt.Printf(student2)
	fmt.Printf("x: %d\n", x)
}
func main() {
	variables()

}
