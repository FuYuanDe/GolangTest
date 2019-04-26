// iface.go
package main

import (
	"fmt"
	"sync"
)

type Session struct {
	name string
	age  int
}

func main() {
	var class sync.Map
	P1 := Session{"Wang", 12}

	class.Store(12, &P1)
	P2, ok := class.Load(12)
	if ok {
		P3, ok := P2.(*Session)
		if ok {
			fmt.Printf("interface assert ok\n")
		} else {
			fmt.Printf("interface assert fail\n")
			return
		}

		fmt.Printf("load ok, %s\n", P3.name)
	} else {
		fmt.Printf("load fail\n")
	}
}
