package main

import (
	"fmt"

	"github.com/JoeSheen/godo/internal/types"
	"github.com/google/uuid"
)

func main() {
	t := types.Task{Id: uuid.New(), Priority: types.Low, Name: "fkjsdfb"}
	fmt.Printf("%v", t)
}
