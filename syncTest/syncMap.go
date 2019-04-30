// syncMap.go
package main

import (
	"fmt"
	"sync"
)

func f(k interface{}, v interface{}) bool {
	fmt.Printf("%v:%v\n", k, v)
	return true
}

func main() {
	var Array sync.Map
	var Group sync.WaitGroup
	Group.Add(2)
	go func() {
		cnt := 1
		for cnt < 10000 {
			Array.Store(cnt, cnt)
			cnt++
		}
		Group.Done()
	}()
	go func() {

		cnt := 1
		for cnt < 10000 {
			Array.Store(cnt, cnt)
			cnt++
		}
		Group.Done()
	}()
	Group.Wait()
	Array.Range(f)

	fmt.Printf("done")
}
