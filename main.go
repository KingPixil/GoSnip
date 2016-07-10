package main

import (
    //"github.com/KingPixil/straw"
    "net/http"
    "os"
    "math/rand"
)

var code string
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
    port := os.Getenv("PORT")
    public := http.FileServer(http.Dir("./static/public/"))
    
    
    http.Handle("/assets", public)
    
    http.HandleFunc("/", landing)
    http.HandleFunc("/new", newHandler)
    
    http.ListenAndServe(":" + port, nil)
}

func landing(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/index.html")
}

func newHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        url := r.FormValue("url")
        
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