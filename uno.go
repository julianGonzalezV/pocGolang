package main
import ("fmt")//fmt Package fmt implements formatted I/O


func forExample() {  
    fmt.Println("forExample")
    
    for i := 1; i <= 5; i++ {
        fmt.Println(i) 
    }
}

//Defer statements
func sample1(){
    fmt.Println("Inside the sample()")
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

    defer sample1()
    fmt.Println("Inside the main()")
}