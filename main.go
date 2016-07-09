package main

import (
    "github.com/KingPixil/straw"
    "string"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    
    http.HandleFunc("/", landing)
    http.HandleFunc("/new", newHandler)
}

func landing(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/index.html")
}

func newHandler(w http.ResponseWriter, r *http.Request) {
    
}