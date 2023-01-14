package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Homepage endpoint hit")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}
