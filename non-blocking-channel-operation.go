package main

import "fmt"
import "time"

func main() {
  messages := make(chan string)

  go func() {
    time.Sleep(time.Second * 3)
    messages <- "blablabla"
  }()

loop:

  for {
    time.Sleep(time.Second * 1)
    select {
      case msg := <-messages:
        fmt.Println("received message", msg)
        break loop
      default:
        fmt.Println("no message received")
    }
  }
}
