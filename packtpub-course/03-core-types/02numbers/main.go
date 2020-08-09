package main

import (
	"fmt"
	"runtime"
)

/// Al alternarse entre tipos de listas podrá observar como es el ahorro en consumo de
/// memoria , asegurese si su data no supera los topes dados con cada tipo de dato
/// y use el adecuado!, le ahorrará problemas a futuro :). claro qu een este caso el
/// deferencial se nota  a partir de 100000 en adelante (si le quita un 0 da igual consumo),
/// es decir que tamibién valide la cantidad de data temporal que va a almacenar
func main() {
	var list []int // El print da TotalAlloc (Consumo de Heap memory) = 403 MiB
	//var list []int8 // increible! da TotalAlloc (Heap) = 54 MiB
	for i := 0; i < 100000; i++ {
		list = append(list, 100)
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("TotalAlloc (Heap) = %v MiB\n", m.TotalAlloc/1024/1024)
}
