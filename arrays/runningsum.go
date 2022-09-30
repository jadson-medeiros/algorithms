package main

func runningSum(nums []int) []int {

	curr_sum := 0

	for i := 0; i < len(nums); i++ {

		curr_sum = curr_sum + nums[i]

		nums[i] = curr_sum
	}

	return nums
}
