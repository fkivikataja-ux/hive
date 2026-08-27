package sprint
/*
1. split the string into words and store them in a string slice.
2. Words are separated by spaces, tabs, and newlines.
*/

func SplitWhitespaces(s string) []string {
	result := []string{}  //this is a slice of string
	str := ""

	for _, r := range s {
		if !(r ==' ' || r== '\t' || r == '\n'  ){ // whitespace definations (space, tab,newline)
			str += string(r)
	}else {
		if str!= ""{
		result = append(result,str) //if str is NOT emty add str to the result
		str  = ""  //emty str
		}
		
			}
	}

	if str != "" {
		result = append(result, str)
	}
	return result

}
