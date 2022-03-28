package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 33
	}()

	fmt.Println(<-ch)

}
