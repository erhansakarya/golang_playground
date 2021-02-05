package main

import "fmt"

// Add adds two integers and return of the sum
func Add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Printf("sum is: %d\n", Add(2, 2))
}
