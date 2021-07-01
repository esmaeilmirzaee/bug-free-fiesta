package task

import (
	"math/rand"
	"time"
)

// Task is a simplate task
type Task struct {
	Id      int
	Title   string
	Content string
}

var tasks []Task

// GenerateId produces a unique identification using random number
func GenerateId() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100000000)
}

// AddTask adds a new task
func AddTask(t Task) bool {
	prevLength := len(tasks)
	tasks = append(tasks, t)
	return len(tasks) > prevLength
}

// GetTasks response with all available tasks
func GetTasks() []Task {
	return tasks
}
