package print
func ReverseAlphabet(step int) string {
	if step <= 0 {
		step = 1
	}
	
	a := ""						   // a variable that contain a string
	for i := 0; i <= 25; i += step {		// start with z, transformed 0, contine until  
		a += string(rune ('z' - i))
	} 
	return a
}