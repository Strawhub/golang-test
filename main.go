package main

import (
	"fmt"
	"os"
	"strconv"
	"test/add"
)

func main() {
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])
	fmt.Println(add.Add(a, b))
}
