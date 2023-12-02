package main

import (
	"dragon/store"
	"fmt"
)

func main() {
	store, err := store.NewDiskStore()
	if err != nil {
		panic(err)
	}
	defer store.Close()

	fmt.Println("Ok la")
}
