package main

import (
	"fmt"
	// "time"
)

func main() {
    // Create a channel named 'messages' for passing strings
    messages := make(chan string)

    // Start a new goroutine
    go func() {
        // Send "ping" into the 'messages' channel
        messages <- "instaview"
    }()

    // Receive the message from the 'messages' channel and assign it to 'msg'
    msg := <-messages

    // Print the received message
    fmt.Println(msg)

	//buffered channel
	message := make(chan string, 2)
	message <- "buffered"
    message <- "channel"

	msg1 := <-message
		fmt.Println(msg1)

	// time.Sleep(time.Second)
	go func() {
        // Send "ping" into the 'messages' channel
        msg1 := <-message
		fmt.Println(msg1)
    }()

	




}

