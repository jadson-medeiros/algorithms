package arrays

// LeetCode - 1351 Count Negative Numbers in a Sorted Matrix

func countNegatives(grid [][]int) (out int) {

	for i, a := range grid {
		for j := range a {
			if grid[i][j] < 0 {
				out++
			}
		}
	}

	return out
}

func countNegatives2(grid [][]int) int {
	out := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] < 0 {
				out++
			}
		}
	}

	return out
}
