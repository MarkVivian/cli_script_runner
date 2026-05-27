package internal

var debugMode bool = true
const destination = "/tmp/"

type JsonData struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Platform string `json:"platform"`
	ScriptURL string `json:"script_url"`
}
