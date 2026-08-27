package sprint

func SortIntegerTable(table []int) []int {
			
	// check every row of the table
	for i:= 0; i < len(table); i++ {  //first for loop look att the row
		for j:= 0; j < len(table)-1-i; j++ {  //second for loop check individual nr
			if table[j]> table[j+1]{
				table[j], table [j+1] = table[j+1], table[j]
			}    // sort the numbers 
		
		}
	
	}
	return table
}