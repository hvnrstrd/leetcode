package main

import "fmt"

// First Unique Character in a String
// Complexity - Time: O(n) - Space: O(1)
// Pattern: Hash Map
// Count character frequencies, then find first with count == 1.

func firstUniqChar(s string) int {
	mapChar := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mapChar[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if mapChar[s[i]] == 1 {
			return i
		}
	}
	return -1
}

// ------------------------------------------------------------------

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
	fmt.Println(firstUniqChar("aabb"))
}
