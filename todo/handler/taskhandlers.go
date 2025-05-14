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

	t.l.Println("method GET request")

	lt := data.GetTasks()

	d, err := json.Marshal(lt)
	if err != nil {
		http.Error(w, "unable to serialize tasks", http.StatusServiceUnavailable)
	}
	w.Write(d)
}

func (t Task) AddTask(w http.ResponseWriter, r *http.Request) {
	return
}
