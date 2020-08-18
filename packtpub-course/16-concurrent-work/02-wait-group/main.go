package main

import (
	"log"
	"sync"
)

func sum(from, to int, wg *sync.WaitGroup, res *int) {
	*res = 0
	for i := from; i <= to; i++ {
		*res += i
	}
	wg.Done()
	return
}

/// Note que con waitGroup ya no es necesario en time.sleep
func main() {
	s1 := 0
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go sum(1, 100, wg, &s1) /// Acá pasamos s1 reference y la actualizamos
	//paso por referencia, no sé si sea recomendado
	wg.Wait()
	log.Println(s1)

}
