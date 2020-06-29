package main

import (
	"fmt"

	"github.com/golang-sample/integers"
)

// Hello return string
func Hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println(integers.Average(10, 20))
	fmt.Println(Hello())
}
