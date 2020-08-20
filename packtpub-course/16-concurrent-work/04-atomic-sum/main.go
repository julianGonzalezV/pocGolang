package main

import (
	"bytes"
	"log"
	"sync"
	"sync/atomic"
	"testing"
)

func main() {
	log.Println("Start main")
	s1 := int64(0)
	wg := &sync.WaitGroup{}
	wg.Add(4) /// Se le indica cuantas goRoutinas va a tener Hilos de ejecución
	go sum(1, 250000000, wg, &s1)
	go sum(250000001, 500000000, wg, &s1)
	go sum(500000001, 750000000, wg, &s1)
	go sum(750000001, 1000000000, wg, &s1)
	wg.Wait()
	log.Println(s1)
	/*2020/08/19 22:33:33 Start main
	2020/08/19 22:33:33 5050*/

}

func sum(from, to int, wg *sync.WaitGroup, res *int64) {
	for i := from; i <= to; i++ {
		atomic.AddInt64(res, int64(i))
	}
	wg.Done()
	return
}

func Test_Main(t *testing.T) {
	for i := 0; i < 10000; i++ {
		var s bytes.Buffer
		log.SetOutput(&s)
		log.SetFlags(0)
		main()
		if s.String() != "5050\n" {
			t.Error(s.String())
		}
	}
}
