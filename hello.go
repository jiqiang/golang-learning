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

  // map
  var numbers1 = make(map[string]int)
  numbers1["one"] = 1
  numbers1["two"] = 2
  numbers1["three"] = 3
  fmt.Println(numbers1) // map[one:1 two:2 three:3]
  fmt.Println(numbers1) // map[two:2 three:3 one:1]
  fmt.Println(numbers1["two"]) // 2
  fmt.Println(len(numbers1)) // 3

  rating := map[string]int {"C": 5, "Go": 4, "Python": 3}
  cRating, ok := rating["C"]
  fmt.Println(cRating) // 5
  fmt.Println(ok) // true
  phpRating, ok := rating["Php"]
  fmt.Println(phpRating) // 0
  fmt.Println(ok) // false
  delete(rating, "C")
  fmt.Println(rating) // map[Go:4 Python:3]

  m1 := make(map[string]string)
  m1["hello"] = "hello"
  fmt.Println(m1) // map[hello:hello]
  m2 := m1
  m2["hello"] = "olleh"
  fmt.Println(m1) // map[hello:olleh]

  // if
  if x := 10; x > 9 {
    fmt.Println("x is greater than 9")
  } else if x == 9 {
    fmt.Println("x is equal to 9")
  } else {
    fmt.Println("x is not greater than 9")
  }

  // for
  sum := 0
  for index := 0; index < 10; index++ {
    sum += index
  }
  fmt.Println("sum is equal to ", sum)

  sum2 := 1
  for sum2 < 10 {
    sum2 += sum2
  }
  fmt.Println("sum is equal to ", sum2)

  // switch
  i := 10
  switch i {
    case 1:
      fmt.Println("i is 1")
    case 2, 3, 4:
      fmt.Println("i is 2 or 3 or 4")
    case 10:
      fmt.Println("i is 10")
      fallthrough
    case 11:
      fmt.Println("i is 11")
    default:
      fmt.Println("i is unknown")
  }
}
