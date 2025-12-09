// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func printsomething(s string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println(s)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	// go printsomething("hello world 1st")

// 	words := []string{
// 		"alphabets",
// 		"beta",
// 		"delta",
// 		"gamma",
// 		"zeta",
// 		"eta",
// 		"theta",
// 		"epsilon",
// 	}
// 	wg.Add(8)
// 	for i, x := range words {
// 		go printsomething(fmt.Sprintf("%d: %s", i, x), &wg)
// 	}
// 	wg.Wait()
// 	// time.Sleep(1 * time.Second)
// 	// printsomething("hello world 2nd")
// }
