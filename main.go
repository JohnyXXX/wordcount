package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		args = args[1:]
		words := strings.Fields(args[0])
		fmt.Println(len(words))
	} else {
		fmt.Println(0)
	}
}
