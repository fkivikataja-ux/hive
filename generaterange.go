package sprint
func GenerateRange(min, max int) []int {  
	if min >= max {
			return nil  		//nil är en nil slice, 
		}

	result := make([]int, max-min)  //medan make ([]int, max-min) är en tom slice med specifik varden , i detta fallet mardet fran max-min.
	
	for i := min; i < max; i++ {
		result [i-min ]= i         // alternativ  append(result, i) instead of result[i-min]+i
	}
	return result
}





