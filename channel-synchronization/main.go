package main

import "fmt"
import "time"

func worker(done chan bool) {
  fmt.Print("working...")
  time.Sleep(time.Second * 3)
  fmt.Println("done")
  done <- true
}

func main() {
  done := make(chan bool, 1)
  go worker(done)
  isDone := <-done // this will block main goroutine execution
  fmt.Println(isDone)
}
