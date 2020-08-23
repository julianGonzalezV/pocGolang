package main

import "log"

/// Nota: Concurrency toma importancia en operaciones que son logicamente independientes
/// como por ejemplo peticiones que recibe un servidor Web, esta peticiones debnerian ser logicamente
/// independientes. Este es un caso práctico para cuando no se deseen que la petición a un server
/// no se vea bloqueada porque otro cliente está procesando una petición que llegó primero
///
/// Otro ejemplo común es cuando se tiene data distriuida en diferentes storage y
/// se requiere reunir toda esa data..se puede tener un GoRoutine or cada storage
/// y combinarlos al final ..Interesante!
func worker(in chan int, out chan int) {
	sum := 0
	for i := range in {
		sum += i
	}
	out <- sum /// sending full result by worker
}

func sum(workers, from, to int) int {
	out := make(chan int, workers) // Setting workers
	in := make(chan int, 4)
	// Loop for runnign the workers needed
	for i := 0; i < workers; i++ {
		go worker(in, out)
	}

	// Sending numbers to be summed to 'in' channel
	for i := from; i <= to; i++ {
		in <- i
	}

	close(in) // Way of telling the parent function (sum(..))
	// that the sum operation has finished

	// gathering information from workers
	sum := 0
	for i := 0; i < workers; i++ {
		sum += <-out
	}

	close(out)
	return sum
}

func main() {
	// note como se va mejerando la version de los pasos previos (ver folders anteriores)
	// ya no se le dice los rangos a procesar!
	res := sum(100, 1, 100)
	log.Println(res)
}
