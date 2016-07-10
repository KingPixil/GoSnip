package main

import (
    //"github.com/KingPixil/straw"
    "net/http"
    "os"
    "math/rand"
    "fmt"
    "log"
)

var code string
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
    port := os.Getenv("PORT")
    i := http.FileServer(http.Dir("static"))
    http.Handle("/", i)
    
    http.HandleFunc("/new", newHandler)
    
    fmt.Println("Listening on Port " + port)
    log.Fatal(http.ListenAndServe(":" + port, nil))
}


func newHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        //url := r.FormValue("url")
        
    }
}

func redir(w http.ResponseWriter, r *http.Request, url string) {
    http.Redirect(w, r, url, 301)
}

func randCode(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}