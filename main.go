package main

import "fmt"

func SomeWork(n int) bool {
	out := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			out <- i
		}
		close(out)
	}()

	for res := range out {
		fmt.Println(res)
	}
	return true
}

func main() {
	// do nothing
	SomeWork(100)
}
