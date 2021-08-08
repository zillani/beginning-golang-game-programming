package main

import "fmt"

func main() {
	result := findGreaterNum(2, 3)
	fmt.Println(result)
}

func findGreaterNum(a int, b int) int {
	if a > b {
		return a
	} else if a < b {
		return b
	} else {
		return 0
	}
}
