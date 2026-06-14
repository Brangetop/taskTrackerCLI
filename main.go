package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"
)

const filename = "tasks.json"

type Task struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command entered")
		os.Exit(1)
	}

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTitle := addCmd.String("title", "", "Task name (required)")
	// addDesc := addCmd.String("desc", "", "Task description")

	doneCmd := flag.NewFlagSet("done", flag.ExitOnError)
	doneTitle := doneCmd.String("title", "", "Completed task name")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])

		if *addTitle == "" {
			fmt.Println("Error: -title flag can't be empty!")
			addCmd.PrintDefaults()
			os.Exit(1)
		}

	case "list":
		listCmd.Parse(os.Args[2:])

		printTasksList()

	case "done":
		doneCmd.Parse(os.Args[2:])

		if *doneTitle == "" {
			fmt.Println("Error: -done flag cant be empty!")
			doneCmd.PrintDefaults()
			os.Exit(1)
		}

	default:
		fmt.Println("Unknown command", os.Args[1])
		os.Exit(1)
	}
}

func loadTasks() ([]Task, error) {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return []Task{}, nil
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error while reading file %w", err)
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing JSON: %w", err)
	}

	return tasks, nil
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Printf("Error while encoding into JSON: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Printf("Error while writiong file: %w", err)
	}

	return nil
}

func printTasksList() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error while loading tasks", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("Task list is empty.")
		return
	}

	fmt.Println("- Task list:")
	for _, task := range tasks {
		statusIcon := "⏳"
		if task.Status == "done" {
			statusIcon = "✅"
		}

		formattedTime := task.CreatedAt.Format("02.01.2006 15:04")

		fmt.Printf("%s [%s] %s\n", statusIcon, task.Status, task.Title)
		if task.Description != "" {
			fmt.Printf("Description: %s\n", task.Description)
		}
		fmt.Printf("Created at:  %s\n", formattedTime)

	}
}

func addTask(title, description string) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error while reading file %v\n", err)
		return
	}

	newTask := Task{
		Title:       title,
		Description: description,
		Status:      "new",
		CreatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	err = saveTasks(tasks)
	if err != nil {
		fmt.Printf("Error while saving new task %v\n", err)
		return
	}

}
