package main

import (
	"ollama/client"
	"fmt"
	"ollama/ollama"
)

const model string = "llama3.2:1b" // for example model
const url string = "http://localhost:11434/api/generate"

func main(){
	go ollama.OllamaServe()
    prompt := ollama.GetPrompt()
	query := ollama.Query{Model: model, Prompt: prompt, Stream: false}
	json_data := client.GetJSON(query)
	response := client.SendData(url, json_data)
	answer := client.GetAnswer(*response)
	text := client.ParseAnswer(answer)
	fmt.Println(text)
}