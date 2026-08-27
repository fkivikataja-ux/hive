package sprint
func ToUpperCase(s string) string {
  result := ""
for _, c := range s {   			//check the hole range of s
	if c >= 'a' && c <= 'z' {		// if c is bigger or equal to 97 and c less or 90
		result += string(c-32)				// c will be c- (differences beetween the )
	}else {
		result += string(c)				//keep iit unchanged if condition not true
	}
	
}
return result
 
}