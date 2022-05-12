package main

import (
    "fmt"
    "os"
    "test/add"
    "strconv"
)

func main() {
    a, _ := strconv.Atoi(os.Args[1])
    b, _ := strconv.Atoi(os.Args[2])
    fmt.Println(add.Add(a, b))
}
