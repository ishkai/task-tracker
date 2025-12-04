package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Write the description of the task.")
			return
		}

		description := strings.Join(os.Args[2:], " ")
		Add(description)
	case "list":
		LoadTask()

		var filter Status

		if len(os.Args) >= 3 {
			switch os.Args[2] {
			case "done":
				filter = StatusDone
			case "todo":
				filter = StatusTodo
			case "in-progress":
				filter = StatusInProgress
			default:
				fmt.Println("Invalid name of status.", os.Args[2])
				return
			}
		}
		List(filter)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Write the description of the task.")
			return
		}

		description := os.Args[2]
		id, err := strconv.Atoi(description)
		if err != nil {
			fmt.Println("ID must be a number!")
			return
		}

		Delete(id)
	case "update":
		if len(os.Args) < 4 {
			return
		}
		idstr := os.Args[2]
		id, _ := strconv.Atoi(idstr)
		description := strings.Join(os.Args[3:], " ")
		Update(id, description)
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Mark the task.")
			return
		}
		description := os.Args[2]
		id, err := strconv.Atoi(description)
		if err != nil {
			fmt.Println("ID must be a number!")
			return
		}
		MarkInProgress(id)
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Mark the task.")
			return
		}
		description := os.Args[2]
		id, err := strconv.Atoi(description)
		if err != nil {
			fmt.Println("ID must be a number!")
			return
		}
		MarkDone(id)
	}

}
