package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"

)

func allUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Users Endpoint Hit")
}