package main

import (
	"fmt"
)

var names = []string{"Jim", "Jane", "Joe", "June"}

func getMapData() map[int]string {
	config := map[int]string{
		101: "1",
		102: "warn",
		103: "1.2.1",
	}
	return config
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// looping over array
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// looping over Map
	map1 := getMapData()

	for key, value := range map1 {
		fmt.Println(key, "=", value)
	}

	// using-break-and-continue-to-control-loops
	for {
		r:=rand.Intn(8)
		if r%3==0 fmt.Print("skipt")
		else fmt.Print("stop") break
	}

}
