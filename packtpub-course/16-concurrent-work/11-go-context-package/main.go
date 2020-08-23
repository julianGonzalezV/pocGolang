package main

import (
	"context"
	"log"
	"time"
)

/*
En situaciones donde se tiene latencia (vía internet por ejemplo), como los
llamados Http. ACÁ LA PREGUNTA ES COMO DEJAMOS De enviar http calls si el server
no responde en un cierto tiempo?
-O como dejar de EJECUTAR una rutina que corre independientemente cuando un evento ocurra?

R/ Una se las soluciones es usando 'context', que es una variable que se pasa
 a través de de una serie de llamados y puede contener en su interior
 valores o estar Vacío (es un contenedor de algo :) ) OJO no se usa para
 enviar valores a diferentes funciones(para eso cree parametros con los tipos del
lenguaje-in, string, bool, struct, etc)

-Se pasa un contexto para recuperar el control de lo que está sucediendo.


*/

/// Este es un ejercicio donde notará como pasar el context a diferentes llamados
//ademas de usarlo para detener la ejecución de un go routine, util para cuando desee
//parar una ejecución compleja y que se tarda más de x tiempo

// OJo acá se cierra el chanel en un go routine,,,evite estas prpacticas, ya que
// es dificil de rastrear donde se cierran los channel y no es muy bien visto

/// countNumbers counts every 100 milliseconds from 0
func countNumbers(c context.Context, r chan int) {
	v := 0
	for {
		select {
		case <-c.Done(): // Is the context done?
			r <- v
			break
		default: // If contextt is not done then keep counting
			time.Sleep(time.Millisecond * 100)
			v++
		}
	}
}

func main() {
	r := make(chan int) // create the channel
	c := context.TODO()

	cl, stop := context.WithCancel(c)
	go countNumbers(cl, r) /// creating count routine

	go func() {
		time.Sleep(time.Millisecond * 100 * 3)
		stop()
	}()

	v := <-r
	log.Println(v)

}
