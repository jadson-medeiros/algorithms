package arrays

// LeetCode - 1672. Richest Customer Wealth

func maximumWealth(accounts [][]int) int {

	wealth := 0

	for i := 0; i < len(accounts); i++ {
		curr_sum := 0

		for j := 0; j < len(accounts[0]); j++ {
			curr_sum = curr_sum + accounts[i][j]
		}

		if curr_sum > wealth {
			wealth = curr_sum
		}
	}

	return wealth
}
