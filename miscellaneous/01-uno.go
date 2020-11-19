package main
import ("fmt")//fmt Package fmt implements formatted I/O
import "time"


func forExample() {  
    fmt.Println("forExample")
    
    for i := 1; i <= 5; i++ {
        fmt.Println(i) 
    }
}

//Defer statements
func sample1(num int){
    fmt.Println("Inside DEFER the sample()",num)
}


func switchExample(inputValue int) {  
    switch inputValue {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two ")
    case 3:
        fmt.Println("tree")
    default:
        fmt.Println("no corresponde")
    }
}

func arrayEg(){
    fmt.Println("::::ARRAYS:::")
    // fIXED sIZE , same type 
    iteger1 := [...] int {1,2,3,4,5}
    //lo mismo pero esta vez indicando el tamanio y no los ...
    iteger2 := [5] int {1,2,3,4,5}

    //Lenght:
    fmt.Println(len(iteger1))
    fmt.Println(len(iteger2))
}

func sliceEg(arrInput [5] string){
    fmt.Println("::::sliceEg:::")
    fmt.Println(arrInput[1:3])
    fmt.Println(len(arrInput[1:3]))
}

//calc is the function name which accepts two integers num1 and num2
//(int, int) says that the function returns two values, both of integer type.
func sumDiff(num1 int, num2 int)(int, int) {  
    fmt.Println("sumDiff")
    sum := num1 + num2
    diff := num1 - num2
    return sum, diff
}

// Esta es una funcion
func displayEmp(emp employee){
    fmt.Println("displayEmp", emp.name)
}

// Convitiendo funcion displayEmp a método 
func(emp employee) displayEmpMet(){
    fmt.Println("displayEmpM", emp.name)
}

type employee struct{
    name string
    address string
    age int
}

func displayNumberLoop(num int) {
	for i:=0; i<=num; i++ {
       // time.Sleep(1 * time.Second)
		fmt.Println("display outside",i)
	}
}

func displayNumberLoopChannel(numCh chan int) {
	for i:=0; i<=4; i++ {
       // time.Sleep(1 * time.Second)
        fmt.Println("channel outside",i)
        
    }
    numCh <- 33 // push data to the channel
}

func main() {
    //declaring a integer variable x
    var x int
    x=3 //assigning x the value 3 
    fmt.Println("x:", x) //prints 3
    
    //declaring a integer variable y with value 20 in a single statement and prints it
    var y int=20
    fmt.Println("y:", y)
    
    //declaring a variable z with value 50 and prints it
    //Here type int is not explicitly mentioned 
    var z=50
    fmt.Println("z:", z)
    
    //Multiple variables are assigned in single line- i with an integer and j with a string
    var i, j = 100,"hello" //aca note que los valores van despues del igual y quiere decir que el primero se llena con 100 osea i  y el segundo con "hello"
    fmt.Println("i and j:", i,j)

    // ejemplo del val en lugar de var, note que se expresa con := los dos puntos antes
    a := 21
	fmt.Println(a)

	/*gives error since a is already declared
	a := 30*/
	fmt.Println(a) 
    
    //::::::::::::CONSTANTS:::::::::::::::::
	const b =10
	fmt.Println(b)
    //    b = 30 no se puede realizar la asignacion por ser constan 
    
     //::::::::::::LOOPS:::::::::::::::::
     forExample()


     var x1 = 50
    if x1 < 10 && (x < 20 || x<25) {
        //Executes if x < 10
        fmt.Println("x is less than 10")
    }else{
        fmt.Println("x is > 10")
    } 

    switchExample(2)

    arrayEg()
    arr := [5] string {"one", "two", "three", "four", "five"}
    sliceEg(arr)
    sm,df := sumDiff(5,6)
    fmt.Println(sm,df)

    //Note la palabra reservada defer, IMPORTANTE: note que así lo coloque de primero dentro del main 
    // de igual forma este no se ejecutará hasta que el MAIN termine
    defer sample1(1)
    fmt.Println("Inside the main()")

    //y que pasaría con multiples Defers?:
    defer sample1(101)
    defer sample1(102)
    defer sample1(103)
    defer sample1(104) //R funcionaría como un stack LIFO lasti int first out

    //Pointers::::::::::
    a1 := 20
    // el oerador & en go se usa para que obtenga la direccion en memoria de una variable
	fmt.Println("Address:",&a1)//imprime la direccoin en memoria de la variable a 
    fmt.Println("Value:",a1)//imprime el valor de la variable a
    // SI! y que es un pointer?
    var b1 *int = &a1//un pointer es tipo de dato que se asigna a una variable en donde se desea la direccion en memoria de otra variable
    // Cómo asi?? 
    // Si, note como la variable b1 es un pointer por el *+Type de tipo int que almacena la direccion de la variable a1
    fmt.Println("Address of pointer b1:",b1)
    // OJO tambien pude hacer var b1  = &a1 (sin *) pero por type inference el se resuelve a *int, sin embargo mejor usarlo para 
    // entender que es un Pointer de manera mas clara   
    fmt.Println("Value of pointer b1:",*b1) // note el * que en este caso se una para obtrener el valor de b1 

    // Structures::::::::::::::: podrian hacer el compare con DTO/POJOS en java
    // a diferencia de Array[] structures permiten elementos del mismo o diferente tipo de datos 
    /*
    type employee struct{
        name string
        address string
        age int
    }*/

    var employee1 employee // instancia
    employee1.name= "julito" // valores
    employee1.address = "andinapolis"
    employee1.age = 32

    
    employee2 := employee{"Raj", "Building-1, Delhi", 25}
    displayEmp(employee1)
    displayEmp(employee2)

    // Methods(not functions) OJO con la anotacion de que no son Functions 
    /*
    Sintaxis
    func (variable variabletype) methodName(parameter1 paramether1type) {  
    }
    */
    // Note que esto no es orientado a objetos, de hecho GO no es orientado a objetos 
    // esto es mas como composicion(queda a futuro cuando sea bien teso, evaluar)
    //"Methods give a feel of what you do in object oriented programs " PERO nOOO  lo es
    employee1.displayEmpMet()
    employee2.displayEmpMet()
    //CONCURRENCY::::::::::::
    /*Reordar que concurrency no es lo mismo que Paralelismo, pues el primero se trata de que los procesos comparte la misma
    CPU y es la CPU es la que gestiona la ejecucion concurrente pero PARALELISMO ES LITERALMENTE 2 TAREAS CORRIENDO AL MISMO 
    TIEMPO, BIEN SEA en otro núcleo o otro nodo fisico/virtual
    */

    //Goroutines: Es la forma en que Go logra ejecutar procesos concurrentemente
    //Goroutine es una funcion que puede correr concurrentemente con otras funciones

    numberLoop := 3 
    /* Note que sin el time.Sleep(1 * time.Second) dentro de displayNumberLoop entonces la funcion gorutine (que se llama con go antes del 
        nombre de la funcion) no se alcanza a ejecutar porque de una sigue con el for en la línea 202 que el programa principal
        y la idea es que el proceso principal que correo no lo tenga que esperar
        > una go rutine precisamente es para eso, para que se pueda continuar la ejcución del método actual(en este caso el Main)
        y en una especied de priceso por debajo - underneath process se ejecute en este caso el displayNumberLoop*/
    go displayNumberLoop(numberLoop)
    //The main() continues without waiting for display()
	for i:=0; i<=numberLoop; i++{
        time.Sleep(1 * time.Second) // sin el time.Sleep(1 * time.Second) acá no se apcanza a ejecutar la go rutine anterior
        fmt.Println("loop in MAIN", i)
    }


    //Channels
    // Note como a comparacion de lo anterior acá no es necesario colocar timers para que se ejecute primero 
    // displayNumberLoopChannel
    nCh := make(chan int)
    go displayNumberLoopChannel(nCh)
    outVar, status := <- nCh //receive data from the channel, espera hasta recibirlos espera == se bloquea 
    //If the status is True it means you received data from the channel. If false, it means you are trying to read from a closed channel
    fmt.Println("Inside main()",status)
	fmt.Println("Printing outVar in main() after taking from channel:",outVar)


}