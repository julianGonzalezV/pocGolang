package future

import (
	"sync"
	"testing"
	"time"
)

func TestStringrError(t *testing.T) {
	future := MaybeString{}
	t.Run("Sucess response => ", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)
		go timeout(t, &wg)
		future.Success(func(s string) {
			t.Fail()
			wg.Done()
		}).Fail(func(e error) {
			t.Log(e.Error())

			wg.Done()
		})
		future.Execute(func() (string, error) {
			return "Execution 1 ", nil
		})
		wg.Wait()
	})
	t.Run("Error response=> ", func(t *testing.T) {

	})
}

func timeout(t *testing.T, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	t.Log("Timeout!")
	t.Fail()
	wg.Done()
}
