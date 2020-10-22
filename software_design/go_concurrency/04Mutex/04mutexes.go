package main

import (
	"sync"
)

// They solve the race condition problem: Two or more goRoutines accessing
// a resource(some memory location) at the same time..then Mutex acts like a
// monitor/parent that allow the permision of accesing or not

// Mutex is used in concurrent programming but in golang we can see another
// mechanism called "channels" BUT it is recommended to meet Mutex cause there
// are scenarios where Channels don't fit well

/// Meter sync.Mutex dentro del struct permitirá lock and unlock de la estructura
/// entera sin necesesidad de estar llamando campo por campo que contenga el struct

/// Nota:
/// Existe el race detector ( run -race ), muy útil para detectar race conditions
/// en nuestr código a la hora de implementar MNutex peero si no usamos Mutex
/// o realizamos un uso incorrecto (bad design) es posible que el race detector
// no logre darse cuenta ejemplo
// 	1) sin usar nada de Mutex
//	2) Colocar el número de procesos en 1 GOMAXPROCS=1 go run -race main.go

type Counter struct {
	sync.Mutex
	value int
}

func counterConcurrent() {
	var wait sync.WaitGroup
	println("Executing counterConcurrent")
	goRoutines := 10
	counter := Counter{}
	wait.Add(goRoutines)
	for i := 0; i < goRoutines; i++ { // 10 routine are lauched
		go func(i int) {
			counter.Lock() // no more Goroutines can access to counter
			counter.value++
			wait.Done()
			defer counter.Unlock() // others can access!!
		}(i)
	}
	wait.Wait()
	counter.Lock()
	defer counter.Unlock()
	println(counter.value)
}
