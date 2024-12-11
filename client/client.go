package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ollama/ollama"
	"time"
)


func GetJSON(query ollama.Query) []byte {
	data, err := json.Marshal(query)
	if err != nil{
		fmt.Println(err)
	}
	return data
}

func SendData(url string, jsonData []byte) *http.Response {
	request, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil{
		fmt.Println(err)
	}
	client := http.Client{Timeout: 10 * time.Second}
	response, err := client.Do(request)
	if err != nil{
		fmt.Println(err)
	}
	return response

}

func GetAnswer(response http.Response) string {
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func ParseAnswer(answer string) string{
	var text map[string]interface{}

	err := json.Unmarshal([]byte(answer), &text)
	if err != nil{
		fmt.Println(err)
	}
	
	value, ok := text["response"]
	if ok{
		strValue, ok := value.(string)
		if ok{
			return strValue
		} else {
			return ""
		}
	} else {
		return ""
	}
}