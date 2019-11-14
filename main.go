package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"Title"` // take note of the required syntax here with the ` key...
	Description string `json:"Desc"`
	Content string `json:"Content"`
}

var Articles []Article

func homePage(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "Hello home page")
	fmt.Println("Loaded home page")
}

func getArticles(writer http.ResponseWriter, r *http.Request) {
	json.NewEncoder(writer).Encode(Articles)
	fmt.Println("Loaded articles")
}

func startServer() {
	port := ":8989"

	// Map handlers to paths
	http.HandleFunc("/articles", getArticles)
	http.HandleFunc("/", homePage)

	error := http.ListenAndServe(port, nil)
	log.Fatal(error)
}

func main() {
	Articles = []Article{
		Article{Title: "Title 1", Description: "Description 1", Content: "Content 1"},
		Article{"Title 2", "Description 2", "Content 2"},
	}

	startServer()
}
