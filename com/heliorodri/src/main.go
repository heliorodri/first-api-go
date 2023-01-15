package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var articles = [4]Article{
	{Id: 1, Title: "Title one", Desc: "Desc one", Content: "Content one"},
	{Id: 2, Title: "Title two", Desc: "Desc two", Content: "Content Two"},
	{Id: 3, Title: "Title three", Desc: "Desc three", Content: "Content three"},
	{Id: 4, Title: "Title four", Desc: "Desc four", Content: "Content four"},
}

func allArticles(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("endpoint hit: All articles")

	json.NewEncoder(writer).Encode(articles)
}

func findArticleById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, error := strconv.Atoi(vars["id"])

	if error != nil {
		panic(error)
	}

	for _, article := range articles {
		if article.Id == id {
			json.NewEncoder(writer).Encode(article)
		}
	}

}

func createArticle(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(body))
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Homepage endpoint hit")
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles)
	myRouter.HandleFunc("/articles", allArticles).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", findArticleById)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Router - API v2 - Mux Router")

	handleRequest()
}
