package pages

import (
	"fmt"
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
	tasks := task.GetTasks()

	pageVariables := PageVariables{
		PageTitle: "[" + fmt.Sprint(len(tasks)) + "]" + "Tasks list | home page",
		Tasks:     tasks,
	}

	t, err := template.ParseFiles("pages/index.html")

	if err != nil {
		log.Fatal("Can't parse index file...", err.Error())
	}

	err = t.Execute(w, pageVariables)

	if err != nil {
		log.Fatal("", err.Error())
	}
}
