
 package main

 import (
  "fmt"
   "runtime"
   "sync"
 )

 // wg is used to wait for the program to finish.
var wg2 sync.WaitGroup

func main(){
	// Averiguar la cantidad de procesadores logicos en la màquina
	defer fmt.Println("Cantidad de procesadores logicos ", runtime.NumCPU())
	// Allocate 1 logical processor for the scheduler to use.
	// Note como al aumentar a la cantidad de Logical processors el tiempo dismunuye :O 
	// es de los beneficios que ofrece OJO NO SIEMPRE PASA! Benchmarking
	runtime.GOMAXPROCS(4)
	// para evitar estar jugando cuantos procesadores logicos entonces de una asigne la capacidad de la máquina 
	//runtime.GOMAXPROCS(runtime.NumCPU())

   // Add a count of two, one for each goroutine.
   	wg2.Add(2)// so las 2 siguinetes goroutines que se van a ejecutar 

	//start := time.Now()	
	//defer elapseTime(start)

   	fmt.Println("Start Goroutines")
	go func() {
		// Schedule the call to Done to tell main we are done.
		defer wg2.Done()
		
		// Display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		 }
	 }()

	 go func() {
		// Schedule the call to Done to tell main we are done.
		defer wg2.Done()
		
		// Display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		 }
	 }()


	
	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg2.Wait()
	
	fmt.Println("\nTerminating Program")
	
}


 
 func printAlphabet(from rune ) {
	// Schedule the call to Done to tell main we are done.
	defer wg2.Done()
	
	// Display the alphabet three times
	for count := 0; count < 3; count++ {
		for char := from; char < from+26; char++ {
			fmt.Printf("%c ", char)
		}
	 }
 }