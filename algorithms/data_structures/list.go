package data_structure

// https://www.geeksforgeeks.org/slices-in-golang/
func Reverse(list []int) []int {
	var s1 []int
	for i := len(list) - 1; i >= 0; i-- {
		s1 = append(s1, list[i])
	}
	return s1
}

func ReverseAbstract(list []interface{}) []interface{} {
	var s1 []interface{}
	for i := len(list) - 1; i >= 0; i-- {
		s1 = append(s1, list[i])
	}
	return s1
}

/// ReverseAbstractSwap uses swap and O(n/2) as so as to improve performance
func ReverseAbstractSwap(list []interface{}) []interface{} {
	s1 := make([]interface{}, len(list))
	copy(s1, list)
	for i := len(list)/2 - 1; i >= 0; i-- {
		aux := len(list) - 1 - i
		//swapping
		s1[i], s1[aux] = s1[aux], s1[i]
	}
	return s1
}
