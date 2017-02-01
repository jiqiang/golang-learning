package main

import "fmt"

import "math/rand"

func main() {

  s1 := rand.NewSource(42)
  r1 := rand.New(s1)

  for i := 0; i < 3; i++ {
    fmt.Print(r1.Intn(100), " ")
  }
  fmt.Println()

  s2 := rand.NewSource(42)
  r2 := rand.New(s2)

  for i := 0; i < 3; i++ {
    fmt.Print(r2.Intn(100), " ")
  }
  fmt.Println()

  s3 := rand.NewSource(43)
  r3 := rand.New(s3)

  for i := 0; i < 3; i++ {
    fmt.Print(r3.Intn(100), " ")
  }
  fmt.Println()
}
