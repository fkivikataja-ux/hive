package sprint
// Programme execution starts here

	
func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	
	if from > to {
		from, to = to, from
	}
	// if either of from or to is less than 0 they should return 0
	if from < 0 {
		from = 0
	}
	
	// if the amount of variable in the arr is bigger 0r less than the arr then make the length to arr 
	if to >= len(arr) {
		to = len(arr) // length start from 1, ndecis start from 0
	}
	// if from > len(arr) {
	// from = len(arr)	
	// }
    
	// (...) is a variadic parameter , takes all the values of the appended values
	arr = append(arr[:from], arr[to:]... ) 
	return arr
}
	