package safe_singleton

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelSingleton(te *testing.T) {
	singleton1 := GetInstance()
	singleton2 := GetInstance()
	n := 5000
	for i := 0; i < n; i++ {
		go singleton1.AddOne()
		go singleton2.AddOne()
	}
	fmt.Printf("El contador es %d \n", singleton1.GetCount())
	var value int
	for value != n*2 {
		value = singleton1.GetCount()
		time.Sleep(10 * time.Millisecond)
	}
	singleton1.Stop()
}

func TestMutexSingleton(t *testing.T) {
	singleton1 := GetInstanceM()
	singleton2 := GetInstanceM()
	n := 5000
	for i := 0; i < n; i++ {
		go singleton1.AddOne()
		go singleton2.AddOne()
	}
	fmt.Printf("El contador es %d \n", singleton1.GetCount())
	var value int
	for value != n*2 {
		value = singleton1.GetCount()
		time.Sleep(10 * time.Millisecond)
	}
}
