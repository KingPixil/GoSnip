package main

import (
	"fmt"
	"log"
	"net/http"
	"math/rand"
)

var url string
var code string
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
    i := http.FileServer(http.Dir("static"))
    
    http.Handle("/", i)
	http.HandleFunc("/new", n)

	log.Println("Listening on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func n(w http.ResponseWriter, r *http.Request)  {
 	if r.Method == "POST" {
 		url = r.FormValue("url")
 		code = randSeq(5)
 		fmt.Fprint(w, code)
 		http.HandleFunc("/" + code, redir)
    }
}

func redir(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, url, 301)
}


func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}