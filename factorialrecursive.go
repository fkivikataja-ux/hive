package sprint

func FactorialRecursive(n int) int {  
	if n <= 1 {
		return 1    // base case
	}
	return n * FactorialRecursive(n-1)  //<-- rercursive case, calls itself
}

