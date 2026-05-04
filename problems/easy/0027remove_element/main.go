package main

import "fmt"

// Remove Element
// Complexity - Time: O(n) - Space: O(1)
// Pattern: Two Pointers
// Left pointer tracks write position, right pointer scans array and skips target value.

func removeElement(nums []int, val int) int {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}

// ------------------------------------------------------------------

func main() {
	nums1 := []int{3, 2, 2, 3}
	fmt.Println(removeElement(nums1, 3))

	nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums2, 2))
}
