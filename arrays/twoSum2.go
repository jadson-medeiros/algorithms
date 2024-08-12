package arrays

import "fmt"

// LeetCode - 167. Two Sum II - Input Array Is Sorted

func twoSum2(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
			curSum := numbers[l] + numbers[r]

			if curSum > target {
					r--
			} else if curSum < target {
					l++
			} else {
					return []int{l + 1, r + 1}
			}
	}

	return []int{}
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	result := twoSum2(numbers, target)
	fmt.Println(result) // Output: [1, 2]
}