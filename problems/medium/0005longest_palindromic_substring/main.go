package main

import "fmt"

// Longest Palindromic Substring
// Complexity - Time: O(n^2) - Space: O(1)
// Pattern: Expand Around Center
// For each character, expand left and right while characters match.
// Two cases: odd length "aba" (center = one char) and even length "abba" (center = two chars).

func longestPalindrome(s string) string {
	best := ""

	expand := func(l, r int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > len(best) {
				best = s[l : r+1]
			}
			l--
			r++
		}
	}

	for i := 0; i < len(s); i++ {
		expand(i, i)
		expand(i, i+1)
	}

	return best
}

// -----------------------------------------------------------------

func main() {
	s := "babadasdasdsddaaadadadadadadd"
	result := longestPalindrome(s)
	fmt.Println(result)
}
