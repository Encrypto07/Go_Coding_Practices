package main

import "fmt"

func main() {

	//creating a channel
	ch := make(chan int)

	//sending 43 to channel
	ch <- 43

	//reciving from the channel
	value := <-ch

	fmt.Println(value)
}
