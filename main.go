package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	baseURL = "https://api.openai.com/v1/chat/completions"
	model   = "gpt-3.5-turbo"
)

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatResponse struct {
	ID      string `json:"id"`
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func main() {

	homedir, _ := os.UserHomeDir()
	err := godotenv.Load(homedir + "/.chatgpt")
	if err != nil {
		log.Fatal(err)
	}

	if os.Getenv("API_KEY") == "" {
		log.Fatal("please check .chatgpt file under home directory and specify API_KEY")
	}

	message := os.Args[1]

	sessionID, err := createChatSession()
	if err != nil {
		fmt.Println("Error creating chat session:", err)
		return
	}

	response, err := interactWithChat(sessionID, message)
	if err != nil {
		fmt.Println("Error interacting with chat:", err)
		return
	}

	// Extract and print the content of the message from the API response
	if len(response.Choices) > 0 {
		fmt.Println(response.Choices[0].Message.Content)
	}
}

func createChatSession() (string, error) {
	requestBody := ChatRequest{
		Model:    model,
		Messages: []ChatMessage{},
	}

	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(requestJSON))
	if err != nil {
		return "", err
	}

	request.Header.Set("Authorization", "Bearer "+os.Getenv("API_KEY"))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	responseJSON, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var chatResponse ChatResponse
	err = json.Unmarshal(responseJSON, &chatResponse)
	if err != nil {
		return "", err
	}

	return chatResponse.ID, nil
}

func interactWithChat(sessionID, message string) (*ChatResponse, error) {
	requestBody := ChatRequest{
		Model: model,
		Messages: []ChatMessage{
			{Role: "user", Content: message},
		},
	}

	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(requestJSON))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+os.Getenv("API_KEY"))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseJSON, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var chatResponse ChatResponse
	err = json.Unmarshal(responseJSON, &chatResponse)
	if err != nil {
		return nil, err
	}

	return &chatResponse, nil
}
