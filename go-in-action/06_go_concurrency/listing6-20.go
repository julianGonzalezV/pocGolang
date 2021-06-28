
 package main

 import (
  "fmt"
   "runtime"
   "Math/rand"
   "sync"
   "time"
 )
// EJEMPLO DE UNBUFFERED CHANNEL 
 var(
	wg sync.WaitGroup
 )

func init(){
	rand.Seed(time.Now().UnixNano())
}


func main(){
	// Averiguar la cantidad de procesadores logicos en la màquina
	defer fmt.Println("Cantidad de procesadores logicos ", runtime.NumCPU())
	
	runtime.GOMAXPROCS(1)
	
	// Así se define un channel SIN BUFFER
	// COURT==FIELD==GROUND
	courtM:= make(chan int)
   	wg.Add(2)// so las 2 siguientes goroutines que se van a ejecutar 



   	fmt.Println("Start Goroutines")
	// Launch two players.
	// Para UNBUFFERED se necesita que ambos jugadores (goRoutines) estén listos en el campo de juego
	go player("Nadal", courtM)
	go player("Djokovic", courtM)
	// hasta esta línea ambas goRoutines están bloqueadas, hasta que se envia una señal (una pelota en éste caso jeje)
	courtM <- 1// se envía el mensaje de inicio


	
	// Wait for the goroutines to finish.
	fmt.Println("Waiting FOR THE Game To Finish")
	wg.Wait()
	
}

func player(name string, court chan int) {
	// Schedule the call to Done to tell main we are done.
	// sin el defer de wg.done saldría error all goroutines are asleep - deadlock!
	defer wg.Done()
	
	// For infinito
	for {
		ball,ok := <- court
		fmt.Printf("ball =  %d ,ok = %t  \n", ball,ok )
		// Si Ok==false, quiere decir que le channel fué cerrado, sino entonces constinua el juego :) 
		if !ok {
			fmt.Printf("El jugador %s ha ganado\n\n\n", name)
			return
		}
		n := rand.Intn(100)
		fmt.Printf("Status del juego para %s decisión %d \n", name, n)
		if n%13 ==0 {
			fmt.Printf("El jugador %s ha perdido\n\n\n", name)
			close(court)
			return
		}
		fmt.Printf("El jugador %s Anota %d punto\n\n\n", name, ball)
		ball++
		// Envìo recuento del partido al channel
		court <- ball

	 }
 }

