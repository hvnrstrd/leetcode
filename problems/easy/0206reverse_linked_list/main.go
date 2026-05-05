package main

import "fmt"

// Reverse Linked List
// Complexity - Time: O(n) - Space: O(1)
// Pattern: Linked List
// Three pointers: prev, curr, next. Flip each arrow then move forward.

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// ------------------------------------------------------------------

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	result := reverseList(head)
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
}
