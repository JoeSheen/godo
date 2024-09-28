/*
Copyright © 2024 Joe Sheen <joe123sheen@hotmail.com>
*/
package main

import (
	"fmt"
	"log"

	"github.com/JoeSheen/godo/cmd"
	"github.com/JoeSheen/godo/internal/sqlite"
)

func main() {
	db, err := sqlite.OpenDBConnection("app.db")
	checkError(err)

	ts, err := db.GetAllTasks()
	checkError(err)

	for _, t := range ts {
		fmt.Printf("%v\n", t)
	}

	t, err := db.GetTaskByID(3)
	checkError(err)
	fmt.Printf("\n\n%v\n\n", t)

	//t, err := db.GetTaskByID(2)
	//if err != nil {
	//log.Fatal(err)
	//}
	//fmt.Printf("%v\n\n", t)
	cmd.Execute()
	/*
		appContext, err := sqlite.OpenDBConnection("app.db")
		if err != nil {
			log.Fatal(err)
		}

		deadline := time.Now().AddDate(1, 0, 0)
		id, err := appContext.CreateTask("test task", "cat1", 5, &deadline)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d", id)
	*/
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
