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
    http.Handle("/static/", http.StripPrefix("/static/", i))
    http.HandleFunc("/", serveTemplate)
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

/* Template serving */
func serveTemplate(w http.ResponseWriter, r *http.Request) {
  lp := path.Join("static", "index.html")
  fp := path.Join("static", r.URL.Path)

  // Return a 404 if the template doesn't exist
  info, err := os.Stat(fp)
  if err != nil {
    if os.IsNotExist(err) {
      http.NotFound(w, r)
      return
    }
  }

  // Return a 404 if the request is for a directory
  if info.IsDir() {
    http.NotFound(w, r)
    return
  }

  tmpl, err := template.ParseFiles(lp, fp)
  if err != nil {
    // Log the detailed error
    log.Println(err.Error())
    // Return a generic "Internal Server Error" message
    http.Error(w, http.StatusText(500), 500)
    return
  }

  if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
    log.Println(err.Error())
    http.Error(w, http.StatusText(500), 500)
  }
}