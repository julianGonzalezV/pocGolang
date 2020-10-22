package main

import (
	"fmt"
	"strings"
	"sync"
)

/// A callback is an anonymous function that will be executed
/// within the context of a different function
/// note la unction que se pasan como parámetros
func toUpperSync(word string, f func(string)) {
	f(strings.ToUpper(word))
}

func toUpperASyncCaller(word string) {
	var wait sync.WaitGroup
	anonymousAsynch := func(v string) {
		fmt.Printf("Result Asynch: %s\n", v)
		wait.Done()
	}
	wait.Add(1)
	toUpperASync(word, anonymousAsynch)
	println("Waiting async response...")
	wait.Wait()
}

/// Asynchronous version
func toUpperASync(word string, f func(string)) {
	go func() {
		f(strings.ToUpper(word))
	}()

}
