Simple CLI Tasks (Go)

Small command-line task manager written in Go to practice basic file I/O and JSON persistence.
Overview

This program stores tasks in a local tasks.json file. Each task has a title, optional description, status, and creation timestamp.
Commands

    Add a task:

go run . add -title "Buy milk" -desc "2 liters"

List tasks:

go run . list

Mark a task done (by title):

    go run . done -title "Buy milk"

Notes

    Tasks are saved in tasks.json next to the program.
    If tasks.json doesn't exist it will be created automatically.
    Titles are used to identify tasks for the done command — use unique titles or extend the program to use IDs.

Requirements

    Go 1.20+ (or recent Go toolchain)
