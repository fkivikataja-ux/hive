package sprint 
/*
1:Capitalize the first letter of each word 
1b: converting the rest of the word to lowercase.
2: A word is defined as a sequence of alphanumeric characters.
Input/outcome:
ToCapitalCase("Hello! Great to see you! How-are-you-doing-2day?")
>> "Hello! Great To See You! How-Are-You-Doing-2day?"
*/
func ToCapitalCase(s string) string {
	result := ""
	IsFirstLetter := true

	for _, c := range s {

		// Är c en bokstav eller siffra?
		IsLetter := (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
		IsNumber := c >= '0' && c <= '9'

		if IsLetter || IsNumber {

			// Första tecknet i ordet → stor bokstav
			if IsFirstLetter {
				if c >= 'a' && c <= 'z' {
					c = c - 32
				}
				IsFirstLetter = false

			} else {
				// Resten av ordet → små bokstäver
				if c >= 'A' && c <= 'Z' {
					c = c + 32
				}
			}

		} else {
			// Inte bokstav/siffra → nästa tecken startar nytt ord
			IsFirstLetter = true
		}

		// Lägg till tecknet i resultatet
		result += string(c)
	}

	return result
}

