package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/joho/godotenv"

	handlers "go-gpt-api/handlers/openAI"
	internal "go-gpt-api/internal/openAI"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	mux := http.NewServeMux()
	openAIClient := internal.NewOpenAI()
	handler := handlers.NewHandler(openAIClient)

	mux.HandleFunc("POST /merge", handler.MergeHandler)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
