package pipeline

import "log"

// ONE OF THE MOST USED: It is useful for designing complex flows of computations where several Goroutines
// are connected as so as to performance a specific logic
// Eg

func LaunchPipeline(amount int) int {
	firstCh := listGenerator(amount)
	secondCh := raisePower(firstCh)
	thirdCh := listSum(secondCh)
	return <-thirdCh
}

func listGenerator(max int) <-chan int {
	outCh := make(chan int, 100)
	go func() {
		for i := 1; i <= max; i++ {
			outCh <- i
			log.Println("listGenerator i= ", i)
		}
		close(outCh)
	}()
	return outCh
}

func raisePower(in <-chan int) <-chan int {
	outCh := make(chan int, 100)
	go func() {
		for dataInt := range in { // receiving data from channel
			log.Println("raisePower i= ", dataInt)
			outCh <- dataInt * dataInt
		}
		close(outCh)
	}()
	return outCh
}

func listSum(in <-chan int) <-chan int {
	outCh := make(chan int, 100)
	go func() {
		var summ int
		for dataInt := range in { // receiving data from channel
			log.Println("listSum i= ", dataInt)
			summ += dataInt
		}
		outCh <- summ
		close(outCh)
	}()
	return outCh
}
