package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Name        string    `json:"title"`
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
	addDesc := addCmd.String("desc", "", "Task description")

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
