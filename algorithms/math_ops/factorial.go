package math_ops

func factorialIter(acc, n int) int {
	if n <= 0 {
		return acc
	}
	return factorialIter(n*acc, n-1)
}

/// Note como se llama recursivamente un anobynous function
/// tambien puede tener una funcion  aparte como lo es factorialIter
func factorial(n int) int {
	var facto func(int, int) int
	facto = func(acc, n int) int {
		if n <= 0 {
			return acc
		}
		return facto(n*acc, n-1)
	}
	return facto(1, n)
}
