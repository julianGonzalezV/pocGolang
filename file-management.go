package main
/*
import "fmt"
import "io/ioutil"
import "os"*/
 // puede usar un solo import con todos los mport  por dentro 
import (
	"fmt"
	"io/ioutil"
	"os"
)

func readFile(name string)(string, error){
	data, err := ioutil.ReadFile(name)
    if err != nil {
        return "",err
	}else{
		return string(data),err
	}
}

func writeOnFile(){
	f, err := os.Create("file1.txt")
	//si existe error creando el archivo entonces lo imprimo y termino la funcion
    if err != nil {
        fmt.Println(err)
        return
	}
	l, err := f.WriteString("Write Line one ")
	if err != nil{
		fmt.Println(err)
        f.Close()
        return
	}
	fmt.Println(l, "bytes written")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}

func main() {  
    d, e := readFile("data1.txt")
	fmt.Println("File reading error", e)
	fmt.Println("Contents of file:", string(d))
	writeOnFile()
}