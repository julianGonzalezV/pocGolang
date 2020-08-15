package main

import (
	"fmt"
	"strings"
)

func main() {
	itemsSold()

	fmt.Println("Index values to Column headers")
	hdr := []string{"empid", "employee", "address", "hours worked", "hourly rate", "manager"}
	csvHdrCol(hdr)
	hdr2 := []string{"EmplOyee   ", "empid", "hours worked", "address", "manager", "hourly rate"}
	csvHdrCol(hdr2)
}

func csvHdrCol(header []string) {
	csvHeadersToColumnIndex := make(map[int]string)
	for i, v := range header {
		//fmt.Println("v ", v)
		v = strings.TrimSpace(v) // removes paces at the end of beggining
		//fmt.Println("v.TrimSpace =>", v)
		switch strings.ToLower(v) {
		case "employee":
			csvHeadersToColumnIndex[i] = v
		case "hours worked":
			csvHeadersToColumnIndex[i] = v
		case "hourly rate":
			csvHeadersToColumnIndex[i] = v
		}
	}
	fmt.Println(csvHeadersToColumnIndex)
}

func itemsSold() {
	items := map[string]int{}
	items["John"] = 41
	items["Celina"] = 109
	items["Micah"] = 24

	for k, v := range items {
		fmt.Printf("%s sold %d items and ", k, v)
		if v < 40 {
			fmt.Println("is below expectations.")
		} else if v > 40 && v <= 100 {
			fmt.Println("meets expectations.")
		} else if v > 100 {
			fmt.Println("exceeded expectations.")
		}
	}
}
