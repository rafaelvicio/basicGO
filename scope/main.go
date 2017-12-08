package main

import (
	"fmt"
)

var y int = 20

func main() {
	x := 10
	fmt.Println(x)
	fmt.Println(z)
}

// erro de escopo
func printx() {
	fmt.Println(x)
}

func printy() {
	fmt.Println(y)
}
