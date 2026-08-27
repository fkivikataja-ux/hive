package sprint
import "fmt"
func Combinations() string {
	result := ""
	first:= true 				//nothing is added yet
	for i:= 0; i <= 9; i++ {
		for j:= i+1; j<= 9; j++ {
			for k:= j+1; k <= 9; k++ {
				if !first {			//if it's not true add comma and space
					result += ", "
				}
				result += fmt.Sprintf("%d%d%d", i, j, k)
				first= false  // change the value into false
			}
			
		}
	}
	return result
}