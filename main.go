package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"myapp/rps"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

const (
	PORT = 8000
)

type User struct {
	name string
	email string
	password string
}

const (
	id = iota
)

func main() {
	// Empty literal construction syntax is {}
	u1:=User{}
	fmt.Printf("%+v\t %v\t %#v\n%T\n%v\n", u1, u1, u1, u1, id)

	i := rand.Intn(1000)
	switch {
	case i != 10:
		log.Println(i)
		fallthrough
	case i > 10:
		log.Println(i)
	case i < 10:
		log.Println(i)
	default:
	}

	myName := "Esmaeil MIRZAEE"
	for index,letter := range myName  {
		if index % 2 == 0 {
			log.Println(index, strings.ToUpper(string(letter)))
		}
	}

	http.HandleFunc("/play", playRound)
	http.HandleFunc("/", homePage)
	log.Println("Server is listening on", PORT)
	if err := http.ListenAndServe(":8000", nil); err !=nil {
		log.Println(err.Error())
	}
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
