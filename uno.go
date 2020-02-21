package main
import ("fmt")//fmt Package fmt implements formatted I/O


func forExample() {  
    fmt.Println("forExample")
    
    for i := 1; i <= 5; i++ {
        fmt.Println(i) 
    }
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
    if x1 < 10 {
        //Executes if x < 10
        fmt.Println("x is less than 10")
    }else{
        fmt.Println("x is > 10")
    } 

    switchExample(2)
}