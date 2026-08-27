package sprint
import "fmt" 				//import need to have the package in quotation mark
 


func Pairs() string {
	result := ""					// defination of var result need to be inside the function
	for i:= 0; i<= 99; i++ {   			//for the first nr starting from 0
		for j:= i+1; j<= 99; j++ {
			if result != ""{  			//  if result is NOt! an emty string
				result += ", "			// ADD to the result comma and empty space
			}
			result += fmt.Sprintf("%02d %02d", i, j) //take result and add the print out in this format for i and j 
		}
	}
	return result

}


