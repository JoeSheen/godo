package main

import (
	"fmt"
	"log"

	"github.com/JoeSheen/godo/internal/sqlite"
	"github.com/JoeSheen/godo/internal/types"
	"github.com/google/uuid"
)

func main() {
	t := types.Task{ID: uuid.New().String(), Priority: types.Low, Title: "fkjsdfb"}
	fmt.Println(t)

	db, err := sqlite.NewSqlDB("database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
