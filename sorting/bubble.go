package sorting

func bubble(arr []int) {
	size := len(arr)
	for i := 0; i < (size - 1); i++ {
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				/* Swapping */
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}
