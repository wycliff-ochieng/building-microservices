package data

import (
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Createdat time.Time `json:"createdat"`
	Completed bool      `json:"completed"`
}

var TaskList = []*Task{
	&Task{
		ID:        2,
		Title:     "Clean my shoes",
		Createdat: time.Now().UTC(),
		Completed: false,
	},
	&Task{
		ID:        4,
		Title:     "Go to church",
		Createdat: time.Now().UTC(),
		Completed: false,
	},
}

func GetTasks() []*Task {
	return TaskList
}

func AddTask(t *Task) {
	t.ID = getNextID()
	TaskList = append(TaskList, t)
}

func getNextID() int {
	lt := TaskList[len(TaskList)-1]
	return lt.ID + 1
}
