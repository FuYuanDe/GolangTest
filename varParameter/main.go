// main.go
package main

import (
	"fmt"
)

func print(strings ...string) {
	fmt.Println("len(strings):", len(strings))
	for value := range strings {
		fmt.Println(value)
	}
}

func main() {
	print("a", "b", "c")
}
