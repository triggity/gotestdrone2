package main

import (
	"fmt"
	"strings"
)

func main() {
	input := []string{"foo", "bar", "thanks"}
	fmt.Printf("Hello World")
	fmt.Printf(strings.Join(input, ", "))

}
