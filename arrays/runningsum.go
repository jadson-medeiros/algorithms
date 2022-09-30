package arrays

// LeetCode 1480 - Running Sum of 1d Array

func runningSum(nums []int) []int {

	curr_sum := 0

	for i := 0; i < len(nums); i++ {

		curr_sum = curr_sum + nums[i]

		nums[i] = curr_sum
	}

	return nums
}
