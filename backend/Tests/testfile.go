//hello world in go
package main

import "fmt"

//write function that calculates the sum of two numbers
func sum(a int, b int) int {
	return a + b
}


func main() {
	fmt.Println("Hello, World! This is a test for go")

	//store input into a and b
	var a, b int
	fmt.Println("Enter two numbers to add")
	fmt.Scanln(&a, &b)
	//print sum of a and b
	fmt.Println("Sum of a and b is", sum(a, b))
	
}
