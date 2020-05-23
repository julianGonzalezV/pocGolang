:::::::::::::COMANDOS IMPORTANTES:::::::::::::::::::::::::::::
go run archivo.go --corre el main del archivo
gofmt -w mutex.go --formatea un archivo (code indentation ) -w ára escribir de nuevo
go fmt path/to/your/package -- formatea todo el paquete
:::::::::::::::::::::::::::::


Grouping of declaration
// Error codes returned by failures to parse an expression.
var (
    ErrInternal      = errors.New("regexp: internal error")
    ErrUnmatchedLpar = errors.New("regexp: unmatched '('")
    ErrUnmatchedRpar = errors.New("regexp: unmatched ')'")
    ...
)

note como se pueden agrupar por conextos, en este caso de erroes
var (
    countLock   sync.Mutex
    inputCount  uint32
    outputCount uint32
    errorCount  uint32
)

//le da quiza ma´s legibilidad al código

::::::::Naming Conventions::::::::::::
Note. the visibility of a name outside a package is determined by whether its first character is upper case:
(lower case, unexported)
(upper case, exported)

:::::::Packages::::::::: minúscula, una sola palabra sin necesidad de guion bajo or camelCase

:::::::Getters: 
owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}

Owner in capital letter is used as Get method (does not require the Get into the name), On the other hand Set requires to be included into the naming convention (see the line 35)

:::::::Interface names: Mtehod name + er 
Reader, Writer, Formatter, CloseNotifier etc.
//Correcto 
if i < f() {
    g()
}
Erróneo (ya que después de f() se generárá en el compilador un semicolon !!! y nos expondremos a comportamientos extraños)
if i < f() 
{
    g()
}

::::Control structures:::::::::wl
