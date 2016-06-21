package main

import (
	"fmt"
	"log"
	"net/http"
	"math/rand"
	"html/template"
)

var url string
var code string
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
    i := http.FileServer(http.Dir("static"))
    http.Handle("/", i))
	http.HandleFunc("/new", n)

	log.Println("Snip is Listening")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func n(w http.ResponseWriter, r *http.Request)  {
 	if r.Method == "POST" {
 		url = r.FormValue("url")
 		code = randCode(5)
 		fmt.Fprint(w, code)
 		http.HandleFunc("/" + code, redir)
    }
}

func redir(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, url, 301)
}


func randCode(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

