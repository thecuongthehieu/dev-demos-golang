package main

import (
	a "dev-demos-golang/packages/subpackage"
	"fmt"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(a.Add(1, 1))
}
