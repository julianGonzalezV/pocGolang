package main

import (
	"fmt"
	"log"
	"net/http"
)

type pageWithCounter struct {
	content string
	heading string
	counter int
}

func (h *pageWithCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("peticion")
	msg := fmt.Sprintf("<h1>%s</h1><br><h2>%s</h2><br>	<h3>cantidad de visitas:  %d </h3>", h.content, h.heading, h.counter+1)
	h.counter = h.counter + 1
	w.Write([]byte(msg))
}

func main() {
	http.Handle("/", &pageWithCounter{content: "Inicio", heading: "Este es el inicio"})
	http.Handle("/pg1", &pageWithCounter{content: "Página 1", heading: "Esta es la Pg1"})
	http.Handle("/pg2", &pageWithCounter{content: "Página 2", heading: "Esta es la Pg2"})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
