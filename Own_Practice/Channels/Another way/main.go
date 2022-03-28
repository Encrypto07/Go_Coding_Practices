package main

import "fmt"

func main() {
	ch := make(chan int, 2) //this is buffered channel, try to not use buffered channel

	ch <- 23
	ch <- 22

	defer fmt.Println(<-ch)
	defer fmt.Println(<-ch)
}
