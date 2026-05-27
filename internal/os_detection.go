package internal

import (
	"runtime"
	"fmt"
)

func DetectOS() string {
	switch os := runtime.GOOS; os {
	case "windows":
		if debugMode {
			fmt.Println("Detected OS: Windows")
		}
		return "Windows"

	case "linux":
		if debugMode {
			fmt.Println("Detected OS: Linux")
		}
		return "Linux"
	
	default:
		if debugMode {
			fmt.Println("Detected OS: Not Supported")
		}
		return "Not Supported"
	}
}