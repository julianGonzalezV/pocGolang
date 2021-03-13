package main

import (
	"fmt"
	"time"
	"log"
)

func sockMerchant(n int, ar []int) int {
	if n <= 0 || n > 100 {
		return 0
	}
	//auxArr := []int{}
	auxMap := map[int]int{}
	countPairs := 0
	for i := 0; i < n; i++ {
		_, exists := auxMap[ar[i]]
		if exists {
			countPairs++
			delete(auxMap, ar[i])
		} else {
			auxMap[ar[i]] = ar[i]
		}
	}
	return countPairs
}

func countingValleys(steps int, path string) int {
	stepCounter := 0
	valleyCounter := 0
	startDown := false
	for _, value := range path {
		//fmt.Println(string(value), " - ", stepCounter)
		if string(value) == "D" && stepCounter == 0 {
			startDown = true
			stepCounter--
		} else if string(value) == "D" {
			stepCounter--
		} else {
			stepCounter++
		}

		if stepCounter == 0 && startDown {
			startDown = false
			valleyCounter++
		}
	}

	return valleyCounter

}

///
func jumpingOnClouds(c []int) int {
	jumpsCounter := 0
	for i := 0; i < len(c); i += 2 {
		//fmt.Println("i=", i, " -  c[i]=", c[i])
		if c[i] == 1 {
			i--
		}
		if c[i] == 0 && i > 0 {
			jumpsCounter++
		}
	}
	if len(c)%2 == 0 {
		jumpsCounter++
	}
	return jumpsCounter
}

///
func repeatedString(s string, n int) int {
	aCounter := 0
	for _, val := range s {
		if string(val) == "a" {
			aCounter++
		}
	}
	aCounter = aCounter * (n / len(s))

	for i := 0; i < (n % len(s)); i++ {
		if string(s[i]) == "a" {
			aCounter++
		}
	}
	return aCounter
}

func hourglassSum(arr [][]int) int {
	fmt.Println(arr[0])
	fmt.Println(arr[0][0])
	fmt.Println(arr[0][1])
	fmt.Println(arr[0][2])
	fmt.Println(arr[0][3])
	for i := 0; i < len(arr[0]); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j])
		}
		fmt.Println("")
	}
	result := 0
	for column := 0; column <= 3; column++ {
		for row := 0; row <= 3; row++ {

			/*fmt.Println(arr[row][column], arr[row][column+1], arr[row][column+2])
			fmt.Println(arr[row+1][column+1])
			fmt.Println(arr[row+2][column], arr[row+2][column+1], arr[row+2][column+2])*/

			summ := arr[row][column] + arr[row][column+1] + arr[row][column+2] +
				arr[row+1][column+1] +
				arr[row+2][column] + arr[row+2][column+1] + arr[row+2][column+2]

			if summ > result {
				result = summ
			}
		}

	}

	return result
}

func leftRotationGoStyle(a []int, d int) []int {
	result := a
	for i := 0; i < d; i++ {

		result = append(result[1:len(result)], result[0])
		fmt.Println("tail =>", result)
	}
	return result
}

func rightRotationGoStyle(a []int, d int) []int {
	result := a
	for i := 0; i < d; i++ {
		head := result[len(result)-1]
		tail := result[0 : len(result)-1]
		result = append([]int{head}, tail...)
		fmt.Println("head =>", head, "tail=> ", tail)
	}
	return result
}

func leftRotationClassicStyle(a []int, d int) []int {
	return nil
}


/// Se convierte la secuencia DNA a una matriz para mejor manipulación
func stringArrToMatriz(arr []string) [][]string{
	dnaMatriz := make([][]string,len(arr))
	for i := 0; i < len(arr); i++ {		
		arrAux:=make([]string,len(arr[i]))
		for j := 0; j < len(arr[i]); j++ {
			arrAux[j] = string(arr[i][j])
		}
		dnaMatriz[i]= arrAux
		
	}
	return dnaMatriz
}

/// Por cada iteracion se recorren filas y columnas 
/// Y las diagonales principales
func searchByRowColumn(arr []string){
	rowString:=""
	columnString:=""
	leftDiagonalString:=""
	rightDiagonalString:=""
	rigthDiagAux:= len(arr)-1
	colAux := 0
	for row := 0; row < len(arr); row++ {	
		leftDiagonalString += string(arr[row][rigthDiagAux])
		rightDiagonalString+= string(arr[row][row])
		rowString = ""
		columnString = ""
		for column := 0; column < len(arr); column++ {
			rowString += string(arr[row][column])
			columnString += string(arr[column][row])
		}
	
		colAux++
		rigthDiagAux--
		fmt.Println("rowString ",rowString)
		fmt.Println("columnString-",columnString)
		fmt.Println("diagonalStringTemp-",leftDiagonalString)
	}
	fmt.Println("diagonalString-",leftDiagonalString)
	fmt.Println("diagonalString-",rightDiagonalString)

}

/// Podría pensarse en hacerse más granular las operaciones con el fin de poder 
/// hacer un uno mejor de goRoutines
func searchByDiagonals(arr []string) {
	rightDiagonalString:=""
    for row := 1-len(arr); row < len(arr)-1; row++{
		rightDiagonalString=""
        for column := 0; column < len(arr); column++ {
            if row == 0 {
				continue;
			}
            if column == 0 {
				continue;
			}

			rightDiagonalString+=string(arr[row-1][column-1])
			rightDiagonalString+=string(arr[row][column])
            //compara con la posicion superior izquierda
			/*
            if (matrix[row - 1][column - 1] == matrix[row][column]) {
				//letterCounter++; else letterCounter = 1;
			

			}*/
            //compara con la posicion superior derecha
            //if (matrix[i + 1][j + 1] == matrix[i][j])letterCounter++; else letterCounter = 1;
            //finaliza como verdadero si hay 4 letras iguales consecutivas
           
			/*if (letterCounter == 4){
				//return true;
			}*/

        }
		fmt.Println("rightDiagonalString => ",rightDiagonalString)
    }
    //return false;
}




func sequenceValidatioin(arr []string) int {
	/*fmt.Println(arr[0])
	fmt.Println(string(arr[0][0]) == "A")
	fmt.Println(string(arr[0][0]))	
	fmt.Println(string(arr[0][1]))
	fmt.Println(arr[0][2])
	fmt.Println(arr[0][3])*/

	
	//fmt.Println("dnaMatriz->",arrStringToMatriz(arr))

	

	/*
	result := 0
	for column := 0; column <= 3; column++ {
		for row := 0; row <= 3; row++ {
			summ := arr[row][column] + arr[row][column+1] + arr[row][column+2] +
				arr[row+1][column+1] +
				arr[row+2][column] + arr[row+2][column+1] + arr[row+2][column+2]

			if summ > result {
				result = summ
			}
		}

	}*/

	start := time.Now()
	elapsed := time.Since(start)

	/*
	start := time.Now()
	dnaMatriz := stringArrToMatriz(arr)	
	fmt.Println("dnaMatriz->",dnaMatriz)
	elapsed := time.Since(start)
	log.Printf("stringArrToMatriz took %s", elapsed)*/

	/*start := time.Now()
	searchByRowColumn2(stringArrToMatriz(arr))
	elapsed := time.Since(start)
	log.Printf("searchByRowColumn2 took %s", elapsed)*/
	
	/*start = time.Now()
	searchByRowColumn(arr)
	elapsed = time.Since(start)
	log.Printf("searchByRowColumn took %s", elapsed)*/

	start = time.Now()
	searchByDiagonals(arr)
	elapsed = time.Since(start)
	log.Printf("searchByDiagonals took %s", elapsed)

	return 0
}

func main() {
	/*fmt.Println(":::::sockMerchant:::::::")
	fmt.Println(sockMerchant(9, []int{10, 20, 20, 10, 10, 30, 50, 10, 20}))
	fmt.Println(sockMerchant(7, []int{1, 2, 1, 2, 1, 3, 2}))

	fmt.Println(":::::countingValleys:::::::")
	fmt.Println(countingValleys(8, "DDUUUUDDDDDUUUUUUDDDDDUUUUDDUUUDDD"))

	fmt.Println(":::::Jumping on the Clouds:::::::")
	fmt.Println(jumpingOnClouds([]int{0, 0, 1, 0, 0, 1, 0}))
	fmt.Println(jumpingOnClouds([]int{0, 1, 0, 0, 0, 1, 0}))

	fmt.Println(":::::repeatedString:::::::")
	fmt.Println(10%3, 50%4)
	fmt.Println("abcac", repeatedString("abcac", 10))
	fmt.Println("aba", repeatedString("aba", 10))
	fmt.Println("a", repeatedString("a", 1000000000000))

	fmt.Println(":::::hourglassSum:::::::")
	fmt.Println("#1 => ", hourglassSum([][]int{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0}}))

	/*fmt.Println("#2 => ", hourglassSum([][]int{
		{-9, -9, -9, 1, 1, 1},
		{0, -9, 0, 4, 3, 2},
		{-9, -9, -9, 1, 2, 3},
		{0, 0, 8, 6, 6, 0},
		{0, 0, 0, -2, 0, 0},
		{0, 0, 1, 2, 4, 0}}))*/

	/*fmt.Println(":::::leftRotatio:::::::")
	fmt.Println(leftRotationGoStyle([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(":::::rightRotationGoStyle:::::::")
	fmt.Println(rightRotationGoStyle([]int{1, 2, 3, 4, 5}, 2))*/

	fmt.Println(":::::Secuence Validation:::::::")
	fmt.Println("#1 => ", sequenceValidatioin([]string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG"}))

}
