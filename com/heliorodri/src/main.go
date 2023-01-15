package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(writer http.ResponseWriter, request *http.Request) {
	articles := Articles{
		Article{Title: "Title one", Desc: "Desc one", Content: "Content one"},
		Article{Title: "Title two", Desc: "Desc two", Content: "Content Two"},
	}

	fmt.Println("endpoint hit: All articles")

	json.NewEncoder(writer).Encode(articles)
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Homepage endpoint hit")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}
