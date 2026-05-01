package main

import (
	"fmt"
	"math"
)

// String to Integer (Atoi)
// Complexity - Time: O(n) - Space: O(1)
// Pattern: String Parsing
// Skip spaces, check sign, read digits one by one and stop on non-digit or overflow.

func myAtoi(s string) int {
	i := 0
	sign := 1
	result := 0

	for i < len(s) && s[i] == ' ' {
		i++
	}

	if i < len(s) && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')
		if result > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		result = result*10 + digit
		i++
	}

	return sign * result
}

// ------------------------------------------------------------------

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("18446744073709551617"))
}
