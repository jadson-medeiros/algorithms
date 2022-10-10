package arrays

// Big O(n)

func sequentialSearch(data []int, value int) bool {

	size := len(data)

	for i := 0; i < size; i++ {
		if data[i] == value {
			return true
		}
	}

	return false
}
