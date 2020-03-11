package main
import "fmt"
import "time"

/*
Channels:::::::
Channels are a way for functions to communicate with each other. It can be thought as a 
medium to where one routine places data and is accessed by another routine.

This subroutine pushes numbers 0 to 9 to the channel and closes the channel*/
// este es el Sender 
func add_to_channel(ch chan int) {	
	fmt.Println("Send data")
	for i:=0; i<10; i++ {
		fmt.Println("Sending",i)
		ch <- i //pushing data to channel
	}
	close(ch) //closing the channel

}

//This subroutine fetches data from the channel and prints it.
// este es el Reader, no cunsummer porque al crear arias go rutine kleyendo del channel no se notifican a todas
func fetch_from_channel(ch chan int) {
	fmt.Println("Read data")
	for {
		//fetch data from channel
x, status := <- ch

		//flag is true if data is received from the channel
//flag is false when the channel is closed
fmt.Println("Read data channe status",status)
if status == true {
			fmt.Println(x)
		}else{
			fmt.Println("Empty channel")
			break	
		}	
	}
}

func fetch2_from_channel(ch chan int) {
	fmt.Println("Read data fetch2")
	for {
		//fetch data from channel
x, status := <- ch

		//flag is true if data is received from the channel
//flag is false when the channel is closed
fmt.Println("fetch2 read data channe status",status)
if status == true {
			fmt.Println(x)
		}else{
			fmt.Println("fetch2 Empty channel")
			break	
		}	
	}
}

func main() {
	//creating a channel variable to transport integer values
	ch := make(chan int)

	//invoking the subroutines to add and fetch from the channel
	//These routines execute simultaneously
	go add_to_channel(ch)
	go fetch_from_channel(ch)
	//go fetch2_from_channel(ch) 
	/* por ahora y desconocimiento al colocar fetch2_from_channel a leer del channel el fetch1 de la linea 61
	no se ejecuta, en el futuro sabré el motivo :)*/

	//delay is to prevent the exiting of main() before goroutines finish
	time.Sleep(5 * time.Second)
	fmt.Println("Inside main()")
}