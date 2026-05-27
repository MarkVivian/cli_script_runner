package main

import (
	"fmt"
	"net"
	"github.com/MarkVivian/cli_script_runner/internal"
)

func checkInternetconnection() bool {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		return false
	}
	conn.Close()
	return true
}


func main() {
	fmt.Println("\t======================================")
	fmt.Println("\t|| Welcome to the CLI Script Runner! ||")
	fmt.Println("\t======================================")


	if !checkInternetconnection() {
		fmt.Println("Error: No internet connection available.")
		return
	}
	
	fmt.Println("\nInternet Connection Secured. You can now run your scripts!")


	internal.DetectOS()
}