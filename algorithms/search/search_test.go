package search

import "testing"

func TestLinear(t *testing.T) {
	result := Linear([]int{1, 2, 3, 5, 4, 6, 7, 8}, 4)
	if !result {
		t.Errorf("Result %t but was expected %t \n", result, true)
	}
}
