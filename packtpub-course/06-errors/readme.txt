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
Guidelines when working with Errors and Panic
Guidelines are just for guidance. They are not set in stone. 
This means, the majority of the time you should follow the guidelines; 
however, there could be exceptions. Some of these guidelines have been
 mentioned previously, but we have consolidated them here for quick reference:

When declaring our own error type, 
the variable needs to start with Err. It should also follow the camel 
case naming convention.
var ErrExampleNotAllowd= errors.New("error example text")

The error string should start with lowercase and not end with punctuation. 
One of the reasons for this guideline is that the error can be returned 
and concatenated with other information relevant to the error.

If a function or method returns an error, it should be evaluated. 
Errors not evaluated can cause the program to not function as expected.

When using panic(), pass an error type as the argument, 
instead of an empty value.

Do not evaluate the string value of an error.

Use the panic() function sparingly.
