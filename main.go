package main

import (
	"log"
	"myapp/rps"
	"net/http"
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
	winner, computerChoice, roundResult := rps.PlayRound(1)
	log.Println(winner, computerChoice, roundResult)
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
