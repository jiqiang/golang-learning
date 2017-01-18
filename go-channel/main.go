package main

import (
    "fmt"
    "runtime"
)

func sum(a []int, c chan int) {
    total := 0
    for _, v := range a {
        total += v
    }
    c <- total // send total to c
}

func fibonacci(n int, c chan int) {
    x, y := 1, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x + y
    }
    close(c)
}

func main() {
    a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    c := make(chan int)

    go sum(a[:len(a)/2], c)
    go sum(a[len(a)/2:], c)

    x, y := <-c, <-c // receive from c

    fmt.Println(x, y, x + y)

    bc := make(chan int, 2)
    bc <- 1
    bc <- 2
    fmt.Println(<-bc)
    fmt.Println(<-bc)

    fmt.Println("channel close")

    c1 := make(chan int, 10)
    go fibonacci(cap(c1), c1)
    for i := range c1 {
        fmt.Println(i)
    }

    fmt.Println("num of cpu:", runtime.NumCPU())
}
