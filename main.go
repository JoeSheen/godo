/*
Copyright © 2024 Joe Sheen <joe123sheen@hotmail.com>
*/
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/JoeSheen/godo/cmd"
	"github.com/JoeSheen/godo/internal/sqlite"
)

func main() {
	db, err := sqlite.OpenDBConnection("app.db")
	checkError(err)

	deadline := time.Now().AddDate(1, 0, 0)
	id, err := db.CreateTask("test task", 3, "category1", &deadline)
	checkError(err)
	fmt.Printf("task with ID: %d created\n", id)

	db.ToggleTaskCompleted(1)
	db.ToggleTaskCompleted(2)

	t, err := db.GetTaskById(2)
	checkError(err)
	fmt.Printf("Task: %v\n", t)
	
	tasks, err := db.GetAllTasksByCompletedStatus(true)
	for _, task := range tasks {
		fmt.Printf("%v\n", task)
	}
	checkError(err)
	cmd.Execute()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
