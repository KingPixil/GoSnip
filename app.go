package main

import (
    "strings"
    "io/ioutil"
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

	log.Println("Snip is Listening")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func n(w http.ResponseWriter, r *http.Request)  {
 	if r.Method == "POST" {
 		url = r.FormValue("url")
 		code = randCode(5)
        temp()
 		http.ServeFile(w, r, "./static/new.html")
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

func temp() {
    	content, err := ioutil.ReadFile("./static/newTemp.html")
        newHTML := strings.Replace(string(content), "{{code}}", code, -1)
        fmt.Println(newHTML)

        err = ioutil.WriteFile("./static/new.html", []byte(newHTML), 0644)
        if err != nil {
            panic(err)
        }
}