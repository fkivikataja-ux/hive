package sprint
/*
You're tasked with creating an iterative 
function that calculates the power of an integer n to the given power. 
Handle negative powers by returning 0.
*/

func ToThePowerIterative(n int, power int) int {

	if power < 0 {  //handle negative powers by returning 0
		return 0
	}
	result := 1
	for i := 0; i < power; i++ { //at iteration 0, loop until amount of power
		result *= n
	
	} 

		return result 
}

