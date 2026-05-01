package main

import "fmt"

// Reverse Integer
// Complexity - Time: O(log n) - Space: O(1)
// Pattern: Math
// Pop digits one by one with % 10 and build the result in reverse order.

func reverse(x int) int {
	var result int
	y := x
	for y != 0 {
		x = y % 10
		result = result*10 + x
		y = y / 10
	}
	if result >= -2147483648 && result <= 2147483647 {
		return result
	}
	return 0
}

// ------------------------------------------------------------------

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(1534236469))
}
