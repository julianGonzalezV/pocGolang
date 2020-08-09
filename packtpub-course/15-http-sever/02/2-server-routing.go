package main

import (
	"log"
	"net/http"
)

type hello struct{}

// ServeHTTP es una función que tendrá el  hello struct (Line 8), el solo
// de tener el método ServeHTTP lo convierte en yn handler, que viene del type
// type Handler interface, intente cambiar el nombre y verá el error en la línea donde
// adicionan el handler hello{}
func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Llamado exitosooo</h1>"
	w.Write([]byte(msg))
}

func main() {

	/// Configurando handlers del server
	/// handler 1: Este ejemplo hace uso de la func HandleFunc de http package
	/// y no recibe un handler como en el ejemplo handler 2
	http.HandleFunc("/chapter1", func(w http.ResponseWriter, r *http.Request) {
		msg := "<h1>Chapter 1</h1>"
		w.Write([]byte(msg))
	})
	/// handler 2 y asi sucesivamente. Note que este ejemplo es usando la func Handle
	/// de http package, use handlers para diseños escalables, el anterior es mu light
	http.Handle("/", hello{})
	/// Se le pasa nil para que tome los handlers anteriores, en lugar del nil, pruebe pasando el handler hello
	/// y notará que cada path (/, /chapter1) siempre esponderá lo mismo que hello()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
