package main

import (
    "fmt"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        fmt.Println(s)
    }
}

func main() {
    go say("world")
    go say("hello")
    say("glenn")

    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}
