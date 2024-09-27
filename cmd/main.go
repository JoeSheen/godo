package main

import (
	"fmt"
	"log"
	"time"

	"github.com/JoeSheen/godo/internal/sqlite"
	"github.com/JoeSheen/godo/internal/types"
)

func main() {
	deadline := time.Now().AddDate(0, 18, 0)
	t := types.Task{Priority: types.Low, Title: "fkjsdfb", Deadline: &deadline}
	fmt.Println(t)

	db, err := sqlite.NewSqlDB("app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
