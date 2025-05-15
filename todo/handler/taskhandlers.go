package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wycliff-ochieng/todo/data"
)

type Task struct {
	l *log.Logger
}

func NewTask(l *log.Logger) *Task {
	return &Task{l}
}

func (t Task) GetTasks(w http.ResponseWriter, r *http.Request) {

	t.l.Println("handling GET request")

	lt := data.GetTasks()

	d, err := json.Marshal(lt)
	if err != nil {
		http.Error(w, "unable to serialize tasks", http.StatusServiceUnavailable)
	}
	w.Write(d)
}

type taskkey struct{}

func (t *Task) AddTask(w http.ResponseWriter, r *http.Request) {

	t.l.Println("handling POST requests")

	tsk := r.Context().Value(taskkey{})
	task, ok := tsk.(data.Task)
	if !ok {
		t.l.Println("task not found in context or tasklist")
		http.Error(w, "invalid task", http.StatusInternalServerError)
	}
	data.AddTask(&task)
}

//handle middleware
