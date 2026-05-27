package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetchScript(repoURL string, dest_suffix string) string {
	// replacing the github.com URL with the raw.githubusercontent.com URL
	rawURL := repoURL
	outputFile := destination + "script." + dest_suffix
	if rawURL == "" {
		fmt.Println("Error: Empty repository URL provided.")
		return ""
	}

	if rawURL[:19] == "https://github.com" {
		rawURL = "https://raw.githubusercontent.com" + rawURL[19:]
	}
	if debugMode {
		fmt.Printf("Fetching script from URL: %s\n", rawURL)
	}

	fmt.Printf("fetching script %v \n", repoURL)

	resp, err := http.Get(rawURL)
	if err != nil {
		fmt.Printf("Error fetching script from %s: %v\n", rawURL, err)
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error fetching script from %s: received status code %d\n", rawURL, resp.StatusCode)
		return ""
	}

	out, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating local file for script from %s: %v\n", rawURL, err)
		return ""
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("Error copying script from %s to %s: %v\n", rawURL, out.Name(), err)
		return ""
	}

	return out.Name() 
}

// todo : add error handling and logging for script execution
// todo : change the parameter to be one specific script instead of a list of scripts.
func windowsScriptExecution(script_id string, script map[string]JsonData) {
	var path string = fetchScript(script[script_id].ScriptURL, "ps1")
	if path == "" {
		fmt.Printf("Failed to fetch script for Windows with ID: %s\n", script_id)
		return
	}
	
	fmt.Printf("found the script %s in location %s ", script_id, path)
}

func linuxScriptExecution(script_id string, script map[string]JsonData) {
	var path string = fetchScript(script[script_id].ScriptURL, "sh")
	if path == "" {
		fmt.Printf("Failed to fetch script for Linux with ID: %s\n", script_id)
		return
	}

	fmt.Printf("found the script %s in location %s ", script_id, path)
}
