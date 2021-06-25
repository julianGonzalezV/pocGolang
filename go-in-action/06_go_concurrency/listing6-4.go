
 package main

 import (
	"time"
  "fmt"
   "runtime"
   "sync"
 )

 // wg is used to wait for the program to finish.
var wg sync.WaitGroup

func main(){
	// Averiguar la cantidad de procesadores logicos en la màquina
	defer fmt.Println("Cantidad de procesadores logicos ", runtime.NumCPU())
	// Allocate 1 logical processor for the scheduler to use.
	// Note como al aumentar a la cantidad de Logical processors el tiempo dismunuye :O 
	// es de los beneficios que ofrece OJO NO SIEMPRE PASA! Benchmarking
	runtime.GOMAXPROCS(1)
	// para evitar estar jugando cuantos procesadores logicos entonces de una asigne la capacidad de la máquina 
	//runtime.GOMAXPROCS(runtime.NumCPU())

   // Add a count of two, one for each goroutine.
   	wg.Add(2)// so las 2 siguinetes goroutines que se van a ejecutar 

	start := time.Now()	
	defer elapseTime(start)

   	fmt.Println("Start Goroutines")
	go printPrime("A")
  	go printPrime("B")
	
	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()
	
	fmt.Println("\nTerminating Program")
	
}

func elapseTime(start time.Time){
	fmt.Println("Tiempo total ", time.Since(start))
}


 // printPrime displays prime numbers for the first 5000 numbers.
 func printPrime(prefix string) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()
	
	next:
	for outer := 2; outer < 50000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
 }