package pages

import (
	"myapp/task"
	"net/http"
)

func AddPage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	t := task.Task{
		Id:      task.GenerateId(),
		Title:   r.FormValue("task_title"),
		Content: r.FormValue("task_content"),
	}

	if ok := task.AddTask(t); ok {
		http.Redirect(w, r, "/", http.StatusCreated)
		// fmt.Fprint(w, "Add a task", t)
	} else {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
	}

}
