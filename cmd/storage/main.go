package main

import (
	"fmt"
	"github.com/kvderevyanko/storage/internal/storage"
	"log"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	restoredFile, restoredErr := st.GetById(file.ID)
	if restoredErr != nil {
		log.Fatal(restoredErr)
	}
	println(restoredFile.Data)

	fmt.Println("Uploaded", file)
}
