package main

import "fmt"

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

func SumAndProduct(a, b int) (int, int) {
  return a + b, a * b
}

func sum1(nums ...int) {
  fmt.Print(nums, " ")
  total := 0
  for _, num := range nums {
    total += num
  }
  fmt.Println(total)
}

func add1(a int) int {
  a = a + 1
  return a
}

func add2(a *int) int {
  *a = *a + 1
  return *a
}

func isOdd(integer int) bool {
  if integer % 2 == 0 {
    return false
  }
  return true
}

func isEven(integer int) bool {
  if integer % 2 == 0 {
    return true
  }
  return false
}

type testInt func(int) bool

func filter(slice []int, f testInt) []int {
  var result []int
  for _, value := range slice {
    if f(value) {
      result = append(result, value)
    }
  }
  return result
}

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

  // single and multiple function returns
  x := 3
  y := 4
  z := 5

  max_xy := max(x, y)
  max_xz := max(x, z)

  fmt.Printf("max(%d, %d) = %d\n", x, y, max_xy)
  fmt.Printf("max(%d, %d) = %d\n", x, z, max_xz)

  x1 := 3
  y1 := 4

  xPLUSy, xTIMESy := SumAndProduct(x1, y1)

  fmt.Printf("%d + %d = %d\n", x1, y1, xPLUSy)
  fmt.Printf("%d * %d = %d\n", x1, y1, xTIMESy)

  // variadic function
  sum1(1, 2)
  sum1(3, 4, 5)

  // pass by value and pointer
  x2 := 3
  y2 := add1(x2)
  fmt.Println(x2)
  fmt.Println(y2)
  x3 := 3
  y3 := add2(&x3)
  fmt.Println(x3)
  fmt.Println(y3)

  // defer
  for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
  }

  // function as type
  slice := []int {1,2,3,4,5,6,7}
  odd := filter(slice, isOdd)
  fmt.Println(odd)
  even := filter(slice, isEven)
  fmt.Println(even)
}
