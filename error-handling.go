package main
import "fmt"
import "os"
import "errors"

//function accepts a filename and tries to open it.
func fileopen(name string) {
    f, er := os.Open(name)

    //er will be nil if the file exists else it returns an error object  
    if er != nil {
        fmt.Println(er)
        return
    }else{
		// si viene un valor entonces hay un error, de esta manera lo atrapamos, estoy seguro que hay maneras más 
		//calidosas por ahora este ejemplo
    	fmt.Println("file opened", f.Name())
    }
}

// Ejemplo de como custimize los errores
func fileopen2(name string)(string, error) {
	f,e := os.Open(name)
	if e != nil{
		return "", errors.New("Error personalizado")
	}else{
		return f.Name(), nil
	}
}

func main() {  
	fileopen("invalid.txt")
	f, e := fileopen2("invalid.txt")
	if e != nil {
        fmt.Println("imprimiendo error ==> ",e)
    }else{
    	fmt.Println("file opened", f)
    }  

}