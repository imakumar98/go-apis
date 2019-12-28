package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"

)




type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}



//let's declare a global Articles array

var Articles []Article

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}


func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAll Articles")
	json.NewEncoder(w).Encode(Articles)
}


func handleRequests() {
	//creates a new instance of mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homeLink)

	myRouter.HandleFunc("/all", returnAllArticles)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}


func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	handleRequests()
}

