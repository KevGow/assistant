package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string  `json:"text"`
		Index        int     `json:"index"`
		Logprobs     *string `json:"logprobs"`
		FinishReason string  `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func main() {
	// Get the message from the command line
	args := os.Args[1:]
	//message := "Can you help me with some coding questions? "
	message := strings.Join(args, " ")

	// Replace YOUR_API_KEY with your OpenAI API key
	apiKey := ""

	// Set the model to use
	model := "text-davinci-003"

	maxTokens := 4000 //2048

	type reqBody struct {
		Model       string `json:"model"`
		Prompt      string `json:"prompt"`
		Temperature int    `json:"temperature"`
		Max_tokens  int    `json:"max_tokens"`
	}

	// Build the request body
	requestBody, err := json.Marshal(reqBody{
		Model:       model,
		Prompt:      message,
		Temperature: 1,
		Max_tokens:  maxTokens,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create an HTTP client
	client := &http.Client{}

	// Build the request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiKey)

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	// Read the response
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var response Response
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("\n********************************************************************")
	fmt.Printf("%s\n", response.Choices[0].Text)
	fmt.Print("\n********************************************************************")
}
