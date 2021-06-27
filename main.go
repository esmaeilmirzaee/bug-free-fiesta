package main

import (

	"log"
	"net/http"
	"text/template"
)

const (
	PORT = 8000
)

func main() {
	http.HandleFunc("/", homePage)
	log.Println("Server is listening on", PORT)
	http.ListenAndServe(":8000", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	t, err:=template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
		return
	}

	err=t.Execute(w,nil)
	
	if err!=nil {
		log.Fatal(err)
		return
	}
}