package arrays

import "fmt"

// LeetCode 217. Contains Duplicate

func containsDuplicate(nums []int) bool {
	setNumbers := make(map[int]struct{})

	for _, num := range nums {
			if _, exists := setNumbers[num]; exists {
					return true
			}
			setNumbers[num] = struct{}{}
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 1}
	fmt.Println(containsDuplicate(nums)) // Output: true
}