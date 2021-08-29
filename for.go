// understanding for loop
package main

import "fmt"

var I int

func main() {
	fmt.Println("For loop 1 - basic loop")
	// regular for loop
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// for loop - without initialization and post
	fmt.Println("For loop 2 - only condtional statement")
	var i int = 1
	for i <= 5 {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()

	// for loop - using functions
	fmt.Println("For loop 3 - using functions")
	for doInit(); doCheck(); doCount() {
		fmt.Printf("%d ", I)
	}
	fmt.Println()
}

func doInit() int {
	I = 1
	return I
}

func doCheck() bool {
	if I <= 5 {
		return true
	} else {
		return false
	}
}

func doCount() {
	I++
}
