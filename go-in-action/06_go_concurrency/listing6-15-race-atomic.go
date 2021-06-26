
 package main

 import (
  "fmt"
  "time"
   "runtime"
   "sync"
   "sync/atomic"
 )

 var(
	shutdown int64
	wg sync.WaitGroup
 )

 /// En este ejemplo usaremos LoadInt64 and StoreInt64, que nos ayuda a leer de forma segura(en cuanto a concurrencia se hable)
 /// para lectura(load) y escritura(StoreInt64) 
func main(){
	runtime.GOMAXPROCS(1)
   	wg.Add(2)// so las 2 siguinetes goroutines que se van a ejecutar 

   	fmt.Println("Start Goroutines")
	go doWork("A")
	go doWork("B")

	/// Simulando espera , en ese tiempo las 2 goroutine anteriores se van a estar ejecutando
	time.Sleep(6 * time.Second)
	// Safely flag it is time to shutdown.
	fmt.Println("Oiga apaguese ")
	/// Escritura concurrentemente segura
	atomic.StoreInt64(&shutdown, 1)

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()
}

func doWork(name string) {
	defer wg.Done()
	/// For sin regla de parada :O forever 
	for {
		fmt.Printf("Ejecutando rutina ---> %s \n", name)
		time.Sleep(1 * time.Second)
		//lectura concurrentemente segura
		///Comente esta parte y notará que nunca termina, ya que esta es la condición de parada
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Me voy a apagar %s \n", name)
			break
		}
	 }
 }

