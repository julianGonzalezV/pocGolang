package safe_singleton

import "fmt"

var addChannel chan bool = make(chan bool)
var getCountChannel chan chan int = make(chan chan int)
var exitChannel chan bool = make(chan bool)

/*type Singleton interface {
	AddOne() int
}*/

type singleton struct {
	count int
}

var instance singleton

func init() {
	fmt.Println("Init will get called on main initialization")
	var count int
	go func(addCh <-chan bool, getCh <-chan chan int, exitCh <-chan bool) {
		for {
			select {
			case <-addChannel:
				count++
			case ch := <-getCountChannel:
				ch <- count
			case <-exitChannel:
				return // end infinite loop started in line10
			}
		}
	}(addChannel, getCountChannel, exitChannel)
}

func GetInstance() *singleton {
	return &instance
}

func (s *singleton) AddOne() {
	addChannel <- true
}

func (s *singleton) GetCount() int {
	countCh := make(chan int) // Note that thi channel is unbuffered, in this case
	// the channel blocks the execution until it receives some data
	defer close(countCh)
	getCountChannel <- countCh
	return <-countCh
}

func (s *singleton) Stop() {
	exitChannel <- true
	close(addChannel)
	close(getCountChannel)
	close(exitChannel)
}
