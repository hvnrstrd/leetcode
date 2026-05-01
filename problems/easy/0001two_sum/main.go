package main

import "fmt"

//Two Sum
//Complexity - Time: O(n) - Space: O(n)
//Pattern: Hashmap
//Use hashmap to store visited numbers.

func twoSum(nums []int, target int) []int {
	mapNums := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, found := mapNums[complement]; found {
			return []int{j, i}
		}
		mapNums[num] = i
	}
	return []int{-1, -1}
}

//-----------------------------------------------------------------

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result[0], result[1])
}
