package main

import (
	"fmt"
	"time"
)

func printsomething(s string) {
	fmt.Println(s)
}

func main() {
	// go printsomething("hello world 1st")

	words := []string{
		"alphabets",
		"beta",
		"delta",
		"gamma",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	for i, x := range words {
		go printsomething(fmt.Sprintf("%d: %s", i, x))
	}

	time.Sleep(1 * time.Second)
	printsomething("hello world 2nd")
}
