package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"update_time"`
}

var tasks []Task

func loadTasks() {
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		return
	}
	json.Unmarshal(file, &tasks)
}

// сохраняем таски и изменения в них
func saveTasks() {
	file, _ := json.MarshalIndent(tasks, "", "    ")
	_ = os.WriteFile("tasks.json", file, 0644)
}

// Добавление тасков
func addTask(text string) {
	task := Task{
		ID:   len(tasks) + 1,
		Text: text, Status: "Pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()}
	tasks = append(tasks, task)
	saveTasks()
}

// апдейт уже сохраненных тасков 
func updateTask(id int, newText string) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Text = newText
			tasks[i].UpdatedAt = time.Now()
			saveTasks()
			return
		}
	}

}

func updateTaskStatus(id int, newStatus string) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = newStatus
			saveTasks()
			return
		}
	}
}

func deleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks()
			return
		}
	}
}

func listTasks(status string) {
	for _, task := range tasks {
		if status == "" || task.Status == status {
			fmt.Printf("[%d] %s -> %s | Created at: %s\n || Updated at: %s\n",
				task.ID, task.Text, task.Status,
				task.CreatedAt.Format("[date:2006-01-02] [time:15:04:05]"),
				task.UpdatedAt.Format("[date:2006-01-02] [time:15:04:05]"))
		}
	}
}

func main() {
	loadTasks()

	command := os.Args[1]
	switch command {
	case "list-done":
		status := "done"
		if len(os.Args) > 2 {
			status = os.Args[2]
		}
		listTasks(status)
	case "update":
		id, _ := strconv.Atoi(os.Args[2])
		newText := os.Args[3]
		updateTask(id, newText)
	case "add":
		text := os.Args[2]
		addTask(text)
	case "delete":
		id, _ := strconv.Atoi(os.Args[2])
		deleteTask(id)
	case "list":
		status := ""
		if len(os.Args) > 2 {
			status = os.Args[2]
		}
		listTasks(status)
	case "done":
		id, _ := strconv.Atoi(os.Args[2])
		updateTaskStatus(id, "Done")
	case "cancel":
		id, _ := strconv.Atoi(os.Args[2])
		updateTaskStatus(id, "Cancel")
	default:
		fmt.Println("Invalid command")
	}
}
