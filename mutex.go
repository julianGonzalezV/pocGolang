/*

MUTEX::::::::
Short from para mutual exclusion
Se usa cuando se desea que un recurso no sea accedido por multiples subrutinas
, sedebe importar el package sync

*/

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

//declare a mutex instance
var mu sync.Mutex
var countm = 0

// count declare count variable, which is accessed by all the routine instances
var count = 0

// process copies count to temp, do some processing(increment) and store back to count
//random delay is added between reading and writing of count variable
func process(n int) {
	//loop incrementing the count by 10
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Int31n(2)) * time.Second)
		temp := count
		temp++ // el temp siempre va a iniciar en más 1  el valor del count
		time.Sleep(time.Duration(rand.Int31n(2)) * time.Second)
		count = temp
	}
	fmt.Println("Count after i="+strconv.Itoa(n)+" Count:", strconv.Itoa(count))
}

/**
Version aplicando mutex que es como un synchronized en java
*/
func processMutex(n int) {
	//loop incrementing the count by 10
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Int31n(2)) * time.Second)
		// Mutex contiene 2 medtodos en lock y unlock con el fin de hacer el recurso(método, variables) exclusivo
		mu.Lock()
		temp := countm
		temp++ // el temp siempre va a iniciar en 1 más que el valor del countm
		time.Sleep(time.Duration(rand.Int31n(2)) * time.Second)
		countm = temp
		mu.Unlock() //debloqueo del recurso
	}
	fmt.Println("countm after i="+strconv.Itoa(n)+" countm:", strconv.Itoa(countm))
}

func main() {
	//loop calling the process() 3 times
	for i := 1; i < 4; i++ {
		go process(i)
		go processMutex(i)
	}

	//delay to wait for the routines to complete
	time.Sleep(25 * time.Second)
	fmt.Println("Final Count:", count) // note que el resultado no es 30 ! :o :o
	/* Se habla de 30 porque que el for de la línea 34 va de 1-->3(included), pero entonces que se hace???
	   asi lo corra varias veces no da los 30 y siempre da valores diferentes
	   "
	   MUTEX AL "RESCATE" !!:*/
	fmt.Println("Final countm:", countm) // acá si es 30! :) :), pero esto es bloqueante ver que otras estrategias hay
}
