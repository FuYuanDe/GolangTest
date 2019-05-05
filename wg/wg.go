// wg.go
package main

import (
	"fmt"
	"sync"
)

func child(wg *sync.WaitGroup, i int) {
	fmt.Printf("child:%d, exit \n", i)
	wg.Done()
}

func main() {

	//定义一个WaitGroup变量
	var wg sync.WaitGroup

	//添加gorouting
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go child(&wg, i)
	}

	//等待gorouting退出
	wg.Wait()
	fmt.Printf("all gorouting exit\n")
	return
}
