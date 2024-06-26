package ch_sync

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func Ch_sync() {

    done := make(chan bool, 1)
    go worker(done)

    <-done
}