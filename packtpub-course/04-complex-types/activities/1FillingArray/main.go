package main

import "fmt"

/// Filling an Array
func activity1(arrIn [10]int) [10]int {
	for i := 0; i < len(arrIn); i++ {
		arrIn[i] = i + 1
	}
	return arrIn
}

func main() {

	fmt.Println(activity1([10]int{}))
}
