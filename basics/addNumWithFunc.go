package main

import "fmt"

//Understand how return statement works.

func main() {
	result := addTwoNum(2, 2)
	fmt.Println(result)
}

func addTwoNum(a int, b int) int {
	return a + b
}