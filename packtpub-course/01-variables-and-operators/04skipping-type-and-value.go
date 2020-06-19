package main

import (
	"fmt"
	"time"
)

// cuando se tenga variables relacionadas a un tema :)
var (
	Debug       bool
	LogLevel    = "info"
	startUpTime = time.Now()
)

func main() {
	fmt.Println(Debug, LogLevel, startUpTime)
}
