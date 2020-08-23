package main

import (
	"fmt"
	"sync"
)

/*
HTTP Servers:
Http server corre como un unico programa y corre sobre el Main routine; luego el mismo
programa se encarga de crear nuevas rutinas por cada cliente que se conecte ..eso lo hace
automáticamente ..LO QUE SI ES COMPLICADO EN PROTEGER EL STATE, POR EJMPLO COMO EVITAR EL
FAMOSO RACE CONDITION ..

TRATE DE VOLVER AL CAPITULO DE SERVERS y asegure que counter sea safety(usando MUTEX)

*/

type Worker struct {
	in, out chan int
	sbw     int         // number of subWorkers
	mtx     *sync.Mutex // mutual exclusion
}

/// readThem es un método que pertenece al struc Worker
func (w *Worker) readThem() {
	w.sbw++
	go func() { // creating a new go routine
		partial := 0
		for i := range w.in { // lopp over 'in' channel
			partial += i
		}
		w.out <- partial
		w.mtx.Lock() // aplico mutual exclusion para gestionar el recurso, del cual requiero que no sea accedido por ninguna otra subrutina
		w.sbw--
		if w.sbw == 0 {
			close(w.out) // a medida que reduce la cantidad de subWorkers(L24) entonces se valida si ya no hay más
			// sbw para cerrar out el channel
		}
		w.mtx.Unlock() //note que el recurso(s) a gestionar en este ejemplo es el w.sbw
		// desbloqueo el mutual exclusion para liberar el recurso...OJO ESTO ES UNA FORMA BLOQUEANTE que puede jugarle en contra, debe gestionarlo muy bien
	}()
}

/// gatherResult es una funcion que retorna el resultado final de todos los sworkers
func (w *Worker) gatherResult() int {
	total := 0
	wg := &sync.WaitGroup{}
	wg.Add(1) // 1 porque solo se generará un solo goRoutine (L39)
	go func() {
		for i := range w.out {
			total += i
		}
		wg.Done() // esperar a que todos los mensajes se procesen
	}()
	wg.Wait() //  esperar a que todos los mensajes se procesen
	return total
}

func main() {
	mtx := &sync.Mutex{}
	in := make(chan int, 100)
	wrNum := 10
	out := make(chan int)
	wrk := Worker{in: in, out: out, mtx: mtx}

	// creating subWorkers
	for i := 1; i <= wrNum; i++ {
		wrk.readThem()
	}

	// Enviar los números a sumar
	for i := 1; i <= 100; i++ {
		in <- i
	}
	close(in) // cerrando el canal para indicar(notifiar) que todo se ha enviado

	//obtener el resultado final
	res := wrk.gatherResult()
	fmt.Println(res)

}
