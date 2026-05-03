package main

import "fmt"

// Palindrome Number
// Complexity - Time: O(log n) - Space: O(1)
// Pattern: Math
// Reverse the number and compare with original.

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var result int
	for y := x; y != 0; y = y / 10 {
		digit := y % 10
		result = result*10 + digit
	}
	return result == x
}

// ------------------------------------------------------------------

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
