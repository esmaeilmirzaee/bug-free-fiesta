package pages

import (
	"fmt"
	"log"
	"net/http"
)

func AddPage(w http.ResponseWriter, r *http.Request) {
	log.Println(r.)
	fmt.Fprint(w, "Add a task")
}
