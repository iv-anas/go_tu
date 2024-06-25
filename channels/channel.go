package main

import "fmt"

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
}
