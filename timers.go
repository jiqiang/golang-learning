package main

import "fmt"
import "time"

func main() {
  t1 := time.NewTimer(time.Second * 2)
  fmt.Println("Timer 1 expired", <-t1.C)

  t2 := time.NewTimer(time.Second * 1)
  go func() {
    fmt.Println("Timer 2 expired", <-t2.C)
  }()
  stop2 := t2.Stop()
  if stop2 {
    fmt.Println("Timer 2 stopped")
  }
}
