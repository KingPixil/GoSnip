package main

import (
    "fmt"
    "strings"
    "io/ioutil"
	"log"
	"net/http"
	"math/rand"
	"os"
)

var url string
var code string
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
    port := os.Getenv("PORT")
    fmt.Println(port)
    i := http.FileServer(http.Dir("static"))
    http.Handle("/", i)
	http.HandleFunc("/new", n)

	log.Println("Snip is Listening")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func n(w http.ResponseWriter, r *http.Request)  {
 	if r.Method == "POST" {
 		url = r.FormValue("url")
 		code = randCode(5)
        fmt.Fprint(w, temp())
 		//http.ServeFile(w, r, "./static/new.html")
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

func temp() string {
    	content, err := ioutil.ReadFile("./static/new.html")
    	if err != nil {
    	    log.Fatal(err)
    	}
        newHTML := strings.Replace(string(content), "{{code}}", code, -1)

        return newHTML
}