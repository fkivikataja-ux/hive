package main 

import (
	"fmt"
	"strings"
	 "os"
	 "bufio"
	 "strconv"
 
)
type DataPoint struct {
	// Id is the nr of the sensor, Key is the name, value gives the value of the measured id
	Id int
	Key string
	Value *float64 //*float64 means Value is a pointer for a float 64 value
} 

func main() {
 // This is the initial state. Every sensor  (ID) starts with Value: nil
 // nil will later be printed as NULL
 
	dataPointArr := []DataPoint{
		{Id: 1, Key: "airTemp", Value: nil},
		{Id: 2, Key: "airPressure", Value: nil},
		{Id: 7, Key: "precipitation", Value: nil},
		{Id: 11, Key: "windSpeed", Value: nil},
		{Id: 12, Key: "windDirection", Value: nil},
		{Id: 13, Key: "humidity", Value: nil},
		{Id: 14, Key: "dewPoint", Value: nil},
		{Id: 15, Key: "soilMoisture", Value: nil},
		{Id: 22, Key: "cloudCover", Value: nil},
	}
 

// Print required start message
fmt.Println("--- Weather Station ---")
// Scanner reads input one line at a time from the terminal.
scanner := bufio.NewScanner(os.Stdin) 

//Keep looping while there are new lines to read
for scanner.Scan(){
// take text from users input
text := scanner.Text() 	
//Remove spaces/newlines around the input.	
text = strings.TrimSpace(text)


//Check with command the user entererd.
switch text {

	case "get":
		//Go through every sensor.
			for _, dataPoint := range dataPointArr{
			//If value is nil, there is no known value.
				if dataPoint.Value == nil {
				fmt.Println(dataPoint.Key + ":NULL")
				} else {
				// *dataPoint.Value gets the actual float stored behind the pointer 
				//strconv.FormatFloat(value, format, precision, bitSize)
				fmt.Println(dataPoint.Key + ":" + strconv.FormatFloat(*dataPoint.Value,'f', -1, 64)) 
				}
		}
	case "clear":
		// Go throough all sensors using their indexes
		for i := range dataPointArr {
		//reset value to nil. nil means NULL
		dataPointArr[i].Value = nil
		}
	case "exit": 
	//Print exit message 
	fmt.Println("Exiting...")
	//Stop main() and stop the program
	return


	default:
		//If the input wasn't get, clear or exit, split the comma, return string with split.
		parts := strings.Split(text, ",")
		// we need exactly 
		//parts[0]= "11"
		//parts[1]= "15.5"

if len(parts) != 2 {
	continue
}
// Convert sensor ID from string to int
//"11" > 11
id, err := strconv.Atoi(parts[0])
//if conversion failed, ignore this input
if err != nil{
	continue
}
//get the value part
valueText := parts[1]
//search through all ID's
	for i := range dataPointArr {
		// check if this is the ID we received?
		if dataPointArr[i].Id == id {
			//the station might send NULL
			if valueText == "NULL" {
				//nil represents NULL
				dataPointArr[i].Value = nil
			}else{
				//convert "15.5" >15.5
				value, err := strconv.ParseFloat(valueText, 64)
				if err != nil {
					break
				}
				//store the address of value
				dataPointArr[i].Value = &value
			}
			//ID found so stop searching
			break
		}
	}
}

}
}

