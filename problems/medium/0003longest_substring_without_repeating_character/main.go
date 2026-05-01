package main

import "fmt"

//longest substring without repeating characters
//Complexity - Time: O(n) - Space: O(min(m, n))
//Pattern: Sliding Window
//Use two pointers to create a sliding window and a hashmap to track characters in the current window.

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int)
	maxLength := 0
	left := 0
	for right := 0; right < len(s); right++ {
		if index, found := charIndex[s[right]]; found && index >= left {
			left = index + 1
		}
		charIndex[s[right]] = right
		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}
	}
	return maxLength
}

//-----------------------------------------------------------------

func main() {
	s := "abcabcbb"
	result := lengthOfLongestSubstring(s)
	fmt.Println(result)
}
