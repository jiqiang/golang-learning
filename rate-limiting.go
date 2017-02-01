package main

import "fmt"
import "time"

func main() {
  // limit by time interval
  requests1 := make(chan int, 5)
  for i := 1; i <= 5; i++ {
    requests1 <- i
  }
  close(requests1)

  limiter1 := time.Tick(time.Millisecond * 500)

  for req := range requests1 {
    <-limiter1
    fmt.Println("request", req, time.Now())
  }

  // limit by number of same time requests
  limiter2 := make(chan time.Time, 3)
  for i := 1; i <= 3; i++ {
    limiter2 <- time.Now()
  }

  go func() {
    for t := range time.Tick(time.Millisecond * 500) {
      limiter2 <- t
    }
  }()

  requests2 := make(chan int, 5)
  for i := 1; i <= 5; i++ {
    requests2 <- i
  }
  close(requests2)

  for req := range requests2 {
    <-limiter2
    fmt.Println("request", req, time.Now())
  }

  // process 3 and wait 1 second
  limiter3 := time.Tick(time.Second * 1)
  requests3 := make(chan int, 30)
  for i := 1; i <= 30; i++ {
    requests3 <- i
  }
  close(requests3)

  index := 1
  for req := range requests3 {
    fmt.Println("request", req)
    if index % 3 == 0 {
      <-limiter3
    }
    index++
  }

}
