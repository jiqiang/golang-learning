package main

import "fmt"

func main() {

  jobs := make(chan int, 5)
  done := make(chan bool)

  go func() {
    for {
      job, more := <-jobs
      if more {
        fmt.Println("received job", job)
      } else {
        fmt.Println("received all jobs")
        done <- true
        return
      }
    }
  }()

  for i := 0; i < 20; i++ {
    jobs <- i
    fmt.Println("send job", i)
  }
  close(jobs)

  <-done

}
