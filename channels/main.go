package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages

	fmt.Println(msg)

	buffer := make(chan string, 2)

	buffer <- "buffered"
	buffer <- "channel"

	fmt.Println(<-buffer)
	fmt.Println(<-buffer)
}
