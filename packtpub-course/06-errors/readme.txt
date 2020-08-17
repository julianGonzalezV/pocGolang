package main

import (
	"errors"
	"fmt"
)

/*
There are three types of errors that you might encounter:

Syntax errors
Runtime errors // during execution
Semantic errors (Incorrect computations)
*/

/// Tips
// Avoid runtime errors
// Para Fors   for i := 0; i <= 10; i++ ...mejor use for i := range nums { ...

// Avoid Semantic errors:

/* When Errors happen, we might want to:
Return the error to the caller
Log the error and continue execution
Stop the execution of the program
Ignore it (this is highly not recommended)
Panic (only in very rare conditions, we will discuss this further later)
*/

/// Acá un ejemplo de como podría crear errores personalizados para una librería http :)
var (
	ErrBodyNotAllowed  = errors.New("http: request method or response status code does not allow body")
	ErrHijacked        = errors.New("http: connection has been hijacked")
	ErrContentLength   = errors.New("http: wrote more than the declared Content-Length")
	ErrWriteAfterFlush = errors.New("unused")
)

func main() {
	ErrBusiness := errors.New("Some bad data")

	fmt.Println("Value => ", ErrBusiness.Error())
	fmt.Printf("BusinessError type: %T", ErrBusiness)

}
