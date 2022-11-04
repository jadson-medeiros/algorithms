package bigo

import "fmt"

func logAllPairsOfArray(array []int) {
	for i := 0; i < len(array); i++ { //O(n)
		for j := 0; j < len(array); j++ { //O(n)
			fmt.Println(array[i], array[j])
		}
	}
}
