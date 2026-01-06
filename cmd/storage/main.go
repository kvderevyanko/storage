package main

import (
	"fmt"
	"github.com/kvderevyanko/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("Hello World", st)
}
