package main

import (
	"fmt"
)

func main() {
	result := sum(20, 20)
	fmt.Printf("The result is: %d\n", result)
}

func sum(a, b int) int {
	return a + b
}
