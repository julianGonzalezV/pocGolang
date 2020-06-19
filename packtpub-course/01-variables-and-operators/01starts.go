package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // First, we have to seed the random number generator
	r := rand.Intn(5) + 1            // number between 0 and 4 plus 1
	stars := strings.Repeat("*", r)
	fmt.Print(stars)
}
