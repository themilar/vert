package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("no args")
	} else {
		for _, x := range os.Args[1:] {
			fmt.Println(x)
		}
	}
}
