package arrays

// LeetCode 1929 - Concatenation of Array

func getConcatenation(nums []int) []int {

	ans := make([]int, 2*len(nums))
	ref_value := 0

	for i := 0; i < len(ans); i++ {
		if ref_value >= len(nums) {
			ref_value = 0
		}

		ans[i] = nums[ref_value]

		ref_value++
	}

	return ans
}
