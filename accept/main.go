package main

import (
	"html/template"
	"net/http"
	"os"
)

type Word struct {
	Word string
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("add.html")
	t.Execute(w, nil)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	word := &Word{Word: r.FormValue("word")}
	t, _ := template.ParseFiles("thanks.html")
	t.Execute(w, word)
}

func main() {
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/add", addHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
