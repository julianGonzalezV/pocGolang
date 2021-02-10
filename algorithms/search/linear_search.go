package search

// Linear search an element in O(n) time
func Linear(inputSlice []int, itemToSearch int) bool {
	for _, v := range inputSlice {
		if v == itemToSearch {
			return true
		}

	}
	return false
}
