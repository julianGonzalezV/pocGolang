package main

import (
	"fmt"
	"time"
	"log"
	"strings"
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
	/*fmt.Println(arr[0])
	fmt.Println(arr[0][0])
	fmt.Println(arr[0][1])
	fmt.Println(arr[0][2])
	fmt.Println(arr[0][3])*/
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

/// searchByRowColumn recorre filas y columnas para evualuar en cada iteracion si existe un dna mutante 
//  Note que se recorren filas y columnas a la misma vez con el fin de optimizar el algoritmo  
func searchByRowsAndColumns(arr []string) bool{
	mutantExists:= false
	rowString:=""
	columnString:=""
	rigthDiagAux:= len(arr)-1
	colAux := 0
	for row := 0; row < len(arr); row++ {	
		rowString = ""
		columnString = ""
		for column := 0; column < len(arr); column++ {
			rowString += string(arr[row][column])
			columnString += string(arr[column][row])
		}
		if mutantDna(rowString){
			mutantExists = true
			break
		}	
		colAux++
		rigthDiagAux--
		/*fmt.Println("rowString ",rowString)
		fmt.Println("isMutant? => ",mutantDna(rowString))
		fmt.Println("columnString-",columnString)
		fmt.Println("isMutant? => ",mutantDna(columnString))*/
	}

	return mutantExists

}

/// searchByLeftDiagonal valida la existencia de un Mutante en la diagonal  con esta inclinacion => \
func searchByLeftDiagonal(arr []string) bool {
	mutantExists:= false
	leftDiagonalString:=""
	rowAux := 0 
	column := 0 
    for diag := 1-len(arr); diag < len(arr)-1; diag++{
		rowAux = 0 
		column = 0 
		leftDiagonalString="" 
		if(diag > 0){
			column = diag 
		}else{
			rowAux = -diag
		}
        for row := rowAux; (row < len(arr) && column < len(arr)) ; row++ {
            //fmt.Println("len(arr)=>",len(arr),"diag => ",diag, "row => ",row,"column => ",column)
			leftDiagonalString += string(arr[row][column])
			column++
        }
		if mutantDna(leftDiagonalString){
			mutantExists = true
			break
		}
		//fmt.Println("leftDiagonalString => ",leftDiagonalString)
		//fmt.Println("isMutant? => ",mutantDna(leftDiagonalString))
    }
    return mutantExists;
}

func searchByRigthDiagonal(arr []string) bool {
	mutantExists:= false
	rightDiagonalString:=""// esla que va en direccion => /
	rowAux := 0 
	column := 0 
	for diag := 1-len(arr); diag < len(arr)-1; diag++{
		rowAux = diag 
		column = 0 
		rightDiagonalString="" 
		if(diag > 0){
			column = diag
			rowAux = len(arr)-1
		}else{
			column = 0 
			rowAux = (len(arr)-1)+diag
		}
        for row := rowAux; row >= 0 && column< len(arr); row-- {
			rightDiagonalString += string(arr[row][column])
			column++
        }
		if mutantDna(rightDiagonalString){
			mutantExists = true
			break
		}
		//fmt.Println("rightDiagonalString => ",rightDiagonalString)
		//fmt.Println("isMutant? => ",mutantDna(rightDiagonalString))
		
    }
    return mutantExists;
}

/// mutantDna contienen todas las secuencias de caracteres que representan un gen mutante 
/// Se decide hacerlo así ya que lo hace más escalable a futuro, imagine si le piden que 
/// la regla de negocio cambió y que ya no es la secuencia de 4 letras iguales :O 
/// Bajo el escenario anterior acá solamente es meter la nueva secuencia o a futuro poderla consultar
/// de alguna base de datos!! 
func mutantDna(dna string) bool {
	mutantDna := false
	correctSequences := []string{"AAAA", "TTTT", "CCCC", "GGGG"}
	for _, seq := range correctSequences{
		/// Si es string que llega a ser evaluado contiene alguna se las secuencias mutantes entonces e
		/// puede trabajar para MAGNETO
		if len(dna) >= len(seq) && strings.Contains(dna, seq){
			//fmt.Println("SE HA ENCONTRATO UNA SECUENCIA MUTANTE! ",dna)
			mutantDna = true
			break
		}
	}
	//fmt.Println("mutantDna return  ",mutantDna)
	return mutantDna
}

/// sequenceValidation Orquesta todos los llamados para validar de DNA
/// Note que se utiliza un OR operator, en caso de que se cumpla uno solo de una retorna el valor 
/// sin tener que ir a los dermás caminos, esto también ayuda al algoritmo
func sequenceValidation(arr []string) bool {
	mutantValid := false
	start := time.Now()	
	if searchByRowsAndColumns(arr) || searchByLeftDiagonal(arr) || searchByRigthDiagonal(arr){
		mutantValid =  true 
	}
	elapsed := time.Since(start)
	log.Printf("sequenceValidatioin tooks %s", elapsed)
	return mutantValid
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

	fmt.Println("#2 => ", hourglassSum([][]int{
		{-9, -9, -9, 1, 1, 1},
		{0, -9, 0, 4, 3, 2},
		{-9, -9, -9, 1, 2, 3},
		{0, 0, 8, 6, 6, 0},
		{0, 0, 0, -2, 0, 0},
		{0, 0, 1, 2, 4, 0}}))
*/
	fmt.Println(":::::leftRotatio:::::::")
	fmt.Println(leftRotationGoStyle([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(":::::rightRotationGoStyle:::::::")
	fmt.Println(rightRotationGoStyle([]int{1, 2, 3, 4, 5}, 2))

	/*fmt.Println(":::::Secuence Validation:::::::")
	fmt.Println("#sequenceValidation es => ", sequenceValidation([]string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG"}))


	fmt.Println(":::::Secuence Validation PEOR ESCENARIO:::::::")
	fmt.Println("#sequenceValidation es => ", sequenceValidation([]string{
		"ATGCGA",
		"CAGTGC",
		"TTATTT",
		"AGACGG",
		"GCGTCA",
		"TCACTG"}))*/

}
