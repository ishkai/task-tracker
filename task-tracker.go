package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func SaveTask() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := json.MarshalIndent(TaskList, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.MkdirAll(filepath.Join(home, "task-tracker"), 0700)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(filepath.Join(home, "task-tracker", "tasklist.json"), data, 0644)
	if err != nil {
		panic(err)
	}
}

func LoadTask() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := os.ReadFile(filepath.Join(home, "task-tracker", "tasklist.json"))
	if err != nil {
		if os.IsNotExist(err) {
			TaskList = []Task{}
			return
		}
		return
	}
	json.Unmarshal(data, &TaskList)
}

func NextId() int {
	maxId := 0
	for _, task := range TaskList {
		if task.Id > maxId {
			maxId = task.Id
		}
	}
	return maxId + 1
}

type Status string

const (
	StatusTodo       = Status("todo")
	StatusInProgress = Status("in_progress")
	StatusDone       = Status("done")
)

var TaskList []Task

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func Add(desc string) {

	LoadTask()

	task := Task{
		Id:          NextId(),
		Description: desc,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	TaskList = append(TaskList, task)
	fmt.Printf("Added new task! ID: %d; Desc: %s\n", task.Id, desc)

	SaveTask()
}

func Delete(id int) {
	LoadTask()
	for index, task := range TaskList {
		if task.Id == id {
			TaskList = append(TaskList[:index], TaskList[index+1:]...)
			SaveTask()
			fmt.Printf("Deleted task! ID: %d; Desc: %s\n", id, task.Description)
			return
		}
	}
	fmt.Println("Task not found!")
}

func List(filter Status) {
	LoadTask()
	for _, task := range TaskList {
		if filter != "" && filter != task.Status {
			continue
		}
		fmt.Printf("%d : %s; status: %s\n", task.Id, task.Description, task.Status)
	}
}

func Update(id int, desc string) {
	LoadTask()

	for i, task := range TaskList {
		if task.Id == id {
			TaskList[i].Description = desc
			TaskList[i].UpdatedAt = time.Now()
			break
		}
	}
	SaveTask()
	fmt.Println("Updated task!", id)
}

func MarkInProgress(id int) {
	LoadTask()

	for i, task := range TaskList {
		if task.Id == id {
			TaskList[i].Status = StatusInProgress
			TaskList[i].UpdatedAt = time.Now()
			break
		}
	}
	SaveTask()
	fmt.Println("Task is in progress!", id)
}

func MarkDone(id int) {
	LoadTask()

	for i, task := range TaskList {
		if task.Id == id {
			TaskList[i].Status = StatusDone
			TaskList[i].UpdatedAt = time.Now()
			break
		}
	}
	SaveTask()
	fmt.Println("Task is done", id)
}
