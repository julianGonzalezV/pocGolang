package main

import (
	"fmt"
	"time"
)

/*SELECT es un switch case de channel where the case selected is the first channel who has executed first all its
computing process
*/

//push data to channel with a 4 second delay
func data1(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "from data1()"
}

//push data to channel with a 2 second delay
func data2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "from data2()"
}

func dataDefault(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from dataDefault()"
}

func main() {
	//creating channel variables for transporting string values
	chan1 := make(chan string)
	chan2 := make(chan string)

	//invoking the subroutines with channel variables
	go data1(chan1)
	go data2(chan2)

	//Both case statements wait for data in the chan1 or chan2.
	//chan2 gets data first since the delay is only 2 sec in data2().
	//So the second case will execute and exits the select block
	select {
	case x := <-chan1:
		fmt.Println(x)
	case y := <-chan2:
		fmt.Println(y)
		/* si descomenta el default: acá se dará cuenta que este el el camino seleccionado, ya que al existir timers en
		ambos channels entonces el default es ejecutado primero, sin importar si incluso tiene un timer time.Sleep(6 * time.Second)
		dentro del defautl,ps ya ingrsó!!
		default:
			fmt.Println("Caso por defecto")*/

	}

}
