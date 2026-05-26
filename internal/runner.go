package internal 

import (
	"fmt"
)

var DebugMode bool = false


func getJsonData(){
	if DebugMode {
		fmt.Println("Fetching JSON data...")
	}

	
}