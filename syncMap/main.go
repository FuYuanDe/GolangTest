// main.go
package main

import (
	"fmt"
	"sync"
)

type Class struct {
	Students sync.Map
}

func handler(key, value interface{}) bool {
	fmt.Printf("Name :%s %s\n", key, value)
	return true
}

func main() {

	class := &Class{}

	//存储值
	class.Students.Store("Zhao", "class 1")
	class.Students.Store("Qian", "class 2")
	class.Students.Store("Sun", "class 3")

	//遍历
	class.Students.Range(handler)

	//查询
	if _, ok := class.Students.Load("Li"); !ok {
		fmt.Println("-->Li not found")
	}

	//查询或者追加
	_, loaded := class.Students.LoadOrStore("Li", "class 4")
	if loaded {
		fmt.Println("-->Load Li success")
	} else {
		fmt.Println("-->Store Li success")
	}

	//删除
	class.Students.Delete("Sun")
	fmt.Println("-->Delete Sun success")

	//遍历
	class.Students.Range(handler)

}
