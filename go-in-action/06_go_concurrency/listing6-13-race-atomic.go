
 package main

 import (
  "fmt"
   "runtime"
   "sync"
   "sync/atomic"
 )

 var(
	counter int64 
	wg sync.WaitGroup
 )

 /// RACE CONDITION: Cuando una o mas go routines acceder descontroladamente el mismo RECURSO
 /// Ejemplo de RACE CONDITION (no lo deseará en ambientes productivos :) )
 /// Solción1:
 ///	haga uso de atomic functions que bloquea(manera tradicional de los demás lenguajes ) 
 ///	el recurso..atomic sirve para tipos NUMÉRICOS y PUNTEROS 
 /// 	note el uso de atomic en incrementCounter y la respuesta es correcta
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



   	fmt.Println("Start Goroutines")

	/// Cpn atomic se incrementa correctamente counter y el resultado final es la suma de la 
	/// cantidad de incrementos por cada go Routine
	go incrementCounter(1)
	go incrementCounter(2)


	
	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()
	
	fmt.Println("Final Counter:", counter)
	
}

func incrementCounter(id int) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()
	
	// Display the alphabet three times
	for count := 0; count < 2; count++ {
		// atomic lo que hace es sincronizar go routines
		atomic.AddInt64(&counter, 1) 
		runtime.Gosched()// Yield(ceder) the thread and be placed back in the queue
	 }
 }

