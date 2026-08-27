/*Create a function that takes a slice of strings and a separator string as its parameters. 
The function should return a new string by concatenating 
all the strings in the slice, 
with each string separated by the provided separator.*/

package sprint

func StrConcatWith(strs []string, sep string) string{
	result := ""

	for i, value := range strs {
		result += value
		
		if i != len(strs)-1{
			result += sep
		} 
	}

	return result
}
