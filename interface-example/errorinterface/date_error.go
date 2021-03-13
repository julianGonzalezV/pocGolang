package errorinterface

import (
	"fmt"
	"time"
)

// DateError ...
type DateError struct {
	Message string
	Date    time.Time
}

// Error() ... Note que al implementar Error ya puede ser tratado como un tipo de la interface error
func (de DateError) Error() string {
	return fmt.Sprintf("%s: %s", de.Date.String(), de.Message)
}

//New ...
func New(message string) DateError {
	return DateError{
		Message: message,
		Date:    time.Now(),
	}
}

//New ...
func NewPointer(message string) *DateError {
	return &DateError{
		Message: message,
		Date:    time.Now(),
	}
}

func OtherF() {
	fmt.Print("HI")
}

//PrintError ... recibe un tipo error o cualquier tipo que cumpla consu firma
func (de DateError) PrintError(err error) string {
	fmt.Println(err.Error())
	return de.Message
}
