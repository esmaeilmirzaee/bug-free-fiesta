package main

import (
	"encoding/json"
	"log"
	"myapp/rps"
	"net/http"
	"strconv"
	"text/template"
)

const (
	PORT = 8000
)

func main() {
	http.HandleFunc("/play", playRound)
	http.HandleFunc("/", homePage)
	log.Println("Server is listening on", PORT)
	http.ListenAndServe(":8000", nil)
}

func playRound(w http.ResponseWriter, r *http.Request) {
	userChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	log.Println(userChoice)
	result := rps.PlayRound(userChoice)
	out, err := json.MarshalIndent(result, "", "")
	if err != nil {
		log.Println(out)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func renderTemplate(w http.ResponseWriter, page string) {
	w.Header().Set("Content-Type", "text/html")
	t, err := template.ParseFiles(page)
	logErrors(err)

	err = t.Execute(w, nil)
	logErrors(err)

}

func logErrors(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}
