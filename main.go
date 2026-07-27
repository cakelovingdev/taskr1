package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const path = "data.json"

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func AddTask() {
	fmt.Println("Added")
	var task Task
	task = Task{
		Id:          1,
		Description: "buy tomato",
		Status:      "todo",
		CreatedAt:   time.Now().Format(time.DateTime),
		UpdatedAt:   time.Now().Format(time.DateTime),
	}
	out, err := json.Marshal(task)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}

func EditTask() {
	fmt.Println("Edited")
}

func DeleteTask() {
	fmt.Println("Deleted")
}

func CreateFile() {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error")
	}
	defer file.Close()
}

func FileExists() bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
}

func main() {
	if FileExists() == true {
		fmt.Println("file exists")
		AddTask()
	} else if FileExists() == false {
		fmt.Println("file doesnt exists")
		CreateFile()
		AddTask()
	}
}
