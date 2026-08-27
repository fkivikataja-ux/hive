package sprint

/*
creating an iterative function that 
calculates the factorial of an integer passed as a parameter.

Make sure to handle errors, 
returning 0 for non-possible values or overflows.
*/


func FactorialIterative(n int) int {
	
	if n <= 0  {
			return 0
		}
	
	for i := 1; i <= n; i++ {  // for i need to start as 1 because result start at 1
	    result *= i
	}
	return result
}



/*
what is factorial?
factorial multiply every number before the factoral number
*/