package barrier

/// Para esperar por más de una respuesta procedende de multiples goRoutines
/// ejemplo: multiples peticiones http para contestar al cliente una sola vez
/// OJO: Es solo un ejemplo pero la definición es clara en decir que es para
/// tareas que se dividen en multiples goRoutines para mejorar el performance
/// de OPERACIONES PARALELAS que se deban contestar en un solo resultado

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var timeoutMilliseconds int = 5000

type barrierResp struct {
	Error error
	Resp  string
}

func barrier(endPoints ...string) {
	nRequest := len(endPoints)
	inCh := make(chan barrierResp, nRequest) /// in es un channel de tipo "barrierResp struct" de buffer igual
	// al tamanio de los endPoint que lleguen a la función
	defer close(inCh)
	responses := make([]barrierResp, nRequest)
	for _, endP := range endPoints {
		go makeRequest(inCh, endP)
	}
	var hasError bool
	for i := 0; i < nRequest; i++ {
		resp := <-inCh
		if resp.Error != nil {
			fmt.Println("ERROR=> ", resp.Error)
			hasError = true
		}
		responses[i] = resp

	}
	if !hasError {
		for _, resp := range responses {
			fmt.Println(resp.Resp)
		}
	}

}
func captureStdout(endPoints ...string) string {
	reader, writer, _ := os.Pipe() // pipe connects io.Writer interface with
	// io.Reader interface so that the reader input is the writer outPut
	os.Stdout = writer /// os.Stdout will be the writer
	out := make(chan string)
	go func() { // go routine that listen while writting on console
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	barrier(endPoints...)
	writer.Close() // Said to the Goroutine that no more input is going to come to i
	temp := <-out  // out channer blocks execution until some value is received from go routine in Line 17
	return temp

}

func makeRequest(out chan<- barrierResp, url string) {
	bResp := barrierResp{}
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeoutMilliseconds) * time.Millisecond),
	}
	httRes, err := client.Get(url)
	if err != nil {
		bResp.Error = err
		out <- bResp
		return
	}
	byt, err := ioutil.ReadAll(httRes.Body)
	if err != nil {
		bResp.Error = err
		out <- bResp
		return
	}

	bResp.Resp = string(byt)
	out <- bResp
}
