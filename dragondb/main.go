package main

import (
	"dragon/store"
	"fmt"
)

func main() {
	s, err := store.New()
	if err != nil {
		panic(err)
	}

	s.Set("foo", "bar")
	s.Set("baz", "qux")
	fmt.Println(s.Get("foo"))
	fmt.Println(s.Get("baz"))
	fmt.Println(s.Get("qux"))
}
