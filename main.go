package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "Hello home page")
	fmt.Print("Loaded home page")
}

func startServer() {
	port := ":8989"

	// Map handlers to paths
	http.HandleFunc("/", homePage)
	error := http.ListenAndServe(port, nil)
	log.Fatal(error)
}

func main() {
	startServer()
}
