package main

import "fmt"

//func main() {
//	forLoop()
//	continueStatement()
//	breakStatement()
//	nestedLoops()
//	rangeKeyword()
//}

func forLoop() {
	/*
		takes up to 3 statements
		statement1 Initializes the loop counter value.
		statement2 Evaluated for each loop iteration. If it evaluates to TRUE, the loop continues. If it evaluates to FALSE, the loop ends.
		statement3 Increases the loop counter value
	*/
	for i := 0; i <= 100; i += 10 {
		fmt.Println(i)
	}
}

func continueStatement() {
	// continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}

func breakStatement() {
	//The break statement is used to break/terminate the loop execution.
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
}

func nestedLoops() {
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}
}

func rangeKeyword() {
	/*
		range keyword is used to more easily iterate over an array, slice or map.
		It returns both the index and the value.
	*/
	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}
}
