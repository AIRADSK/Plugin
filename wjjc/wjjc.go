package main

import (
	"fmt"
)

var (
	read  string
	write string
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Scanln(&write)
	fmt.Scanln("%s", &read)
	fmt.Println(read)
}
