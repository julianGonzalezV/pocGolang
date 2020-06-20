package main

import (
	"fmt"
	"math/rand"
	"time"
)

// cuando se tenga variables relacionadas a un tema :)
// note que se le omite el value ...colocando false por defecto
var (
	Debug       bool
	LogLevel    = "info"
	startUpTime = time.Now()
)

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	fmt.Println(Debug, LogLevel, startUpTime)

	//Type Inference Gone Wrong, si se quita el int64 falla
	var seed int64 = 1234456789
	rand.Seed(seed)

	//Short variable declaration
	Debug := true
	LogLevel := "info"
	startUpTime := time.Now()
	fmt.Println(Debug, LogLevel, startUpTime)

	Debug = false
	fmt.Println(Debug, LogLevel, startUpTime)

	//Declaring Multiple Variables from a Function
	Debug2, LogLevel2, startUpTime2 := getConfig()
	fmt.Println("getConfig => ", Debug2, LogLevel2, startUpTime2)

	// Type only
	var start, middle, end float32
	fmt.Println(start, middle, end)
	// Initial value mixed type
	var name, left, right, top, bottom = "one", 1, 1.5, 2, 2.5
	fmt.Println(name, left, right, top, bottom)
	// works with functions also
	var x, y, z = getConfig()
	fmt.Println(x, y, z)

	// Non-English Variable Names
	デバッグ := false
	日志级别 := "info"
	ይጀምሩ := time.Now()
	_A1_Μείγμα := ""
	fmt.Println("Non-English vars =>", デバッグ, 日志级别, ይጀምሩ, _A1_Μείγμα)

}
