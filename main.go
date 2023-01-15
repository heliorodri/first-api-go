package main

import (
	"encoding/json"
	"fmt"
	"io"
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

type Articles []Article

var articles = Articles{
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
	body, _ := io.ReadAll(r.Body)

	var article Article

	json.Unmarshal(body, &article)

	articles = append(articles, article)

	json.NewEncoder(w).Encode(article)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var article Article

	json.Unmarshal(body, &article)

	if article.Id <= 0 {
		article.Id = id
	}

	for index, articleToUpdate := range articles {
		if articleToUpdate.Id == id {
			articles[index] = article
		}
	}

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, article := range articles {
		if article.Id == id {
			articles = append(articles[:index], articles[index+1:]...)
		}
	}
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Homepage endpoint hit")
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", createArticle).Methods("POST")
	myRouter.HandleFunc("/articles", allArticles)
	myRouter.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/articles/{id}", findArticleById)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Router - API v2 - Mux Router")

	handleRequest()
}
