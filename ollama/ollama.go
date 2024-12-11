package ollama

import (
	"fmt"
	"os/exec"
)


type Query struct {
    Model string
    Prompt string
    Stream bool
}

func GetPrompt() string{
    var prompt string
    fmt.Println("Type a question:")
    fmt.Scan(&prompt)
    return prompt
}


func OllamaServe(){
    // Start local ollama server 
    cmd := exec.Command("ollama", "serve")
    cmd.Run()
}


