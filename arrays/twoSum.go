package arrays

import "fmt"

// LeetCode - 1. Two Sum

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for i, num := range nums {
			complement := target - num

			if index, found := numsMap[complement]; found {
					return []int{index, i}
			}

			numsMap[num] = i
	}

	return []int{-1, -1}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result) // Output: [0, 1]
}