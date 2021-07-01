package main

import (
	"log"
	"myapp/pages"
	"net/http"
)

func main() {
	http.HandleFunc("/", pages.Home)
	http.HandleFunc("/add", pages.AddPage)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err.Error())
	} else {
		log.Println("Server is listening on localhost:8000.")
	}
}
