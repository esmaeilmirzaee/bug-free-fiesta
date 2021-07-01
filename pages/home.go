package pages

import (
	"log"
	"myapp/task"
	"net/http"
	"text/template"
)

type PageVariables struct {
	PageTitle string
	Tasks     []task.Task
}

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("pages/index.html")

	if err != nil {
		log.Fatal("Can't parse index file...", err.Error())
	}

	err = t.Execute(w, nil)

	if err != nil {
		log.Fatal("", err.Error())
	}

}
