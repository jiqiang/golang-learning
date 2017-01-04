package main

import "fmt"

func main() {
  fmt.Println("Hello, world!")

  // array
  ar := [10]int {10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

  // slice
  var a, b, c, d, e, f, g []int
  a = ar[2:5]
  b = ar[3:5]
  c = ar[:8] // ar[0:8]
  d = ar[8:] // ar[8:len(ar)]
  e = ar[:] // ar[0:len(ar)]
  fmt.Println(a) // [30 40 50]
  fmt.Println(b) // [40 50]
  fmt.Println(c) // [10 20 30 40 50 60 70 80]
  fmt.Println(d) // [90 100]
  fmt.Println(e) // [10 20 30 40 50 60 70 80 90 100]

  f = a
  a[0] = 101
  fmt.Println(f) // [101 40 50]

  g = ar[2:5]
  fmt.Println(len(g)) // 3
  fmt.Println(cap(g)) // 8

  g = append(g, 110)
  fmt.Println(g) // [101 40 50 110]
  fmt.Println(ar) // [10 20 101 40 50 110 70 80 90 100]

  arr1 := []int {1, 2, 3}
  arr2 := make([]int, len(arr1))
  copy(arr2, arr1)
  fmt.Println(arr2) // [1 2 3]
}
