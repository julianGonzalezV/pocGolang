package main

import (
	"fmt"
	"sync"
)

/// Here we use the standar Done() function //allway recommended
func WaitG1() {
	fmt.Println("::::::::::::Start WaitG1:::::::")
	var wait sync.WaitGroup
	wait.Add(1) // Set the WaitGroup counter in one, if you set 2 or 3 then the following lines must be 2-3 go routines
	go func() {
		fmt.Println("WaitG1 Done")
		wait.Done() // Send done messajge to wait group as so as  release the locked waiting
		//  Done decrements the WaitGroup counter by one
		// si no se libera entonces fatal error: all goroutines are asleep - deadlock!
	}()
	wait.Wait() // wait till done message is sent
}

/// Hear we decremente waitgroup counter(wait.Add(-1) not recommended), clud
/// exploit it
func WaitG2() {
	fmt.Println("::::::::::::Start WaitG2:::::::")
	var wait sync.WaitGroup
	wait.Add(1) // Set the WaitGroup counter in one, if you set 2 or 3 then the following lines must be 2-3 go routines
	go func() {
		fmt.Println("WaitG2 Done")
		wait.Add(-1) // Send done messajge to wait group as so as  release the locked waiting
		//  Done decrements the WaitGroup counter by one
		// si no se libera entonces fatal error: all goroutines are asleep - deadlock!
	}()
	wait.Wait() // wait till done message is sent
}

/// WaitG3: Example If we know in advance how many Goroutines we are going to launch, we can
/// also call the Add method just once
func WaitG3() {
	fmt.Println("::::::::::::Start WaitG3:::::::")
	var wait sync.WaitGroup
	goRoutines := 5
	wait.Add(goRoutines)

	for i := 0; i <= goRoutines; i++ {
		go func(message int) {
			fmt.Printf("WaitG3 go routine %d  Done \n", message)
			wait.Done()
		}(i)
	}
	wait.Wait()
}
