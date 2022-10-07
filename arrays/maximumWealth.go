package arrays

import "fmt"

// LeetCode - 1672. Richest Customer Wealth

func maximumWealth(accounts [][]int) int {

	wealth := 0

	for i := 0; i < len(accounts); i++ {
		curr_sum := 0

		for j := 0; j < len(accounts[0]); j++ {
			curr_sum += accounts[i][j]
		}

		if curr_sum > wealth {
			wealth = curr_sum
		}
	}

	return wealth
}

// Other Solution

func maximumWealth2(accounts [][]int) int {
	total := len(accounts)

	fmt.Println(total)
	wealth := 0

	for _, customer := range accounts {
		sum_value := sumBank(customer)

		if sum_value > wealth {
			wealth = sum_value
		}
	}

	return wealth
}

func sumBank(banks []int) int {
	sum := 0

	for _, bank := range banks {
		sum += bank
	}

	return sum
}
