
 package main

 import (
  "fmt"
   "runtime"
   "Math/rand"
   "sync"
   "time"
 )
// Ejemplo de BUFFERED Channel 
 const(
	numberGoroutines = 4
	taskLoad = 10 // cantidad de trabajo a realizar
 )

 var(
	wgbf sync.WaitGroup
 )

 /// Init se llama antes que el man, por lo regular se usa para configuraciones
func init(){
	// Number random generator
	rand.Seed(time.Now().Unix())
}


func main(){
	// Averiguar la cantidad de procesadores logicos en la màquina
	defer fmt.Println("Cantidad de procesadores logicos ", runtime.NumCPU())
	
	runtime.GOMAXPROCS(1)
	
	//Se crea el channel con un buffer que pueda o soporte la cantidad de trabajo a realizar
	tasksCh:= make(chan string, taskLoad)
	wgbf.Add(numberGoroutines)// se le indica las rutinas que aguantará
	//Se crean las goRoutines dinámicamente 
	for gr := 1; gr <= numberGoroutines; gr++{
		go worker(gr,tasksCh)
	}
	for toDo := 1; toDo <= taskLoad; toDo++{
		tasksCh <-  fmt.Sprintf("Task #%d",toDo)
	}

   	fmt.Println("Start Goroutines")	
	   // A pesar de que se le dice que cierre el channel realmente no se cierra de inmediato
	   // sino que el main thread espera a que se culmine 
	close(tasksCh)/// Despues de cerrar se puede continuar leyend del buffe hasta que se acabe peero NO ENVIAR MENSAJES
	
	// Wait for the goroutines to finish.
	fmt.Println(":::::::Closed channel BUT the main is Waiting FOR THE Job To Finish")
	wgbf.Wait()
	
}

func worker(trabajador int, tareas chan string) {
	// Schedule the call to Done to tell main we are done.
	// sin el defer de wg.done saldría error all goroutines are asleep - deadlock!
	defer wgbf.Done()
	
	// For infinito
	for {
		// Se espera por la tareas entrantes para realizar por el worker
		tarea,ok := <- tareas
		fmt.Printf("tarea =  %s ,ok = %t  \n", tarea,ok )
		// Si Ok==false, quiere decir que el BUFFERED está VACÍO o el channel fué cerrado , sino entonces continua 
		if !ok {
			fmt.Printf("El Trabajador %d se va  descansar\n\n\n", trabajador)
			return
		}
		fmt.Printf("El Trabajador: %d : Inicia tarea %s\n", trabajador, tarea)
		sleep := rand.Intn(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)// Simulando la duración de la tarea

		fmt.Printf("El trabajador %d ha terminado la tarea %s \n", trabajador, tarea)
		

	 }
 }

