package main

import (
	"fmt"
	"sync"
)

// Cuanodo salga Error  "fatal error: all goroutines are asleep - deadlock!"
//Entonces revisar:
// 1) No se establece el buffer o se establece pero no con la cantidad suficiente (cantidad de items, ñike buckets)
// 		y se procede an enviar data si tener en cuenta este tamaño
//2) Se lee del channel más de lo que le llega(tamaño del buffer) ejemplo digo que el buffer es 2 y tengo 3 prints
// esperando la data ..deberá tener 2 prints!
// 3) Doy un buffer X y no cumplo con ese tamaño y envío un valor menos(relacionado con la causa #1)
// ejemplo: Doy u buffer de 3 pero hago envío solo de 2 mensajes, debería enviar 3!

// Para entenderlo mejor, lo que realmente pasa es que el goRoutine se bloquea al llenarse su buffer(con el tamaño que le demos)

func bufferedChannel() {
	fmt.Println("Start bufferedChannel")
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch <- 2 //fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//fmt.Println(<-ch) //fatal error: all goroutines are asleep - deadlock!, el tamaño es 2!
}

func unBufferedChannel() {
	fmt.Println("Start unBufferedChannel")
	ch := make(chan int)
	go readThem(ch)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	//.....ch <- n
}
func readThem(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func unBufferedChannelV2() {
	fmt.Println("Start unBufferedChannelV2")
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ch := make(chan int)
	go readThemV2(ch, wg) // new routine
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	ch <- 6
	//.....ch <- n

	close(ch) // el channel se cierra en el Main routine
	//(recuerde que estamos en el main porque este es el que llamó la función)
	// hasta que todo se envía
	// Sino lo cerramos entonces obtendremos un all goroutines are asleep - deadlock!
	// aunque se cierre el channel los mensajes se siguen enviando al readThemV2 routine
	// y es porque el channel se cierra hasta que que el receptor de los mensajes los reciba todos.
	wg.Wait()
}

func readThemV2(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println(i)
	}
	wg.Done() // se asegura que todo se completa
}

func main() {
	bufferedChannel()

	// Solución más eficaz :), peero aveces no muestra todos los mensajes cuando el computo de cada mensaje que se envíe
	// se realice muy rápido :(
	unBufferedChannel()

	// Solución al problema anterior: El waitGroup al rescate!!
	unBufferedChannelV2()
}
