package main

import (
	"context"
	"fmt"
	"time"
)

func withCancel1() {
	//建立一个父context
	parentCtx, cancel := context.WithCancel(context.Background())
	go func() {
		//从父context衍生出子context
		childCtx, _ := context.WithCancel(parentCtx)
		go func() {
			//子协程运行
			for {
				select {
				case <-childCtx.Done():
					fmt.Printf("childCtx Done, Child exit\n")
					return
				case <-time.After(time.Second * 2):
					fmt.Println("Child is running...")
				}
			}
		}()

		//父协程运行
		for {
			select {
			case <-parentCtx.Done():
				fmt.Println("parentCtx Done, Parent exit")
				return
			case <-time.After(time.Second * 6):
				fmt.Println("parent is running...")
			}
		}
	}()

	time.Sleep(8 * time.Second)

	fmt.Printf("Time Up, Close parentCtx\n")
	//取消父context
	cancel()
	time.Sleep(4 * time.Second)
}

func withCancel2() {
	//定义一个空的ctx
	rootCtx := context.TODO()

	//从根ctx衍生子ctx
	ctx, cancel := context.WithCancel(rootCtx)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("cancel called, gorouting exit\n")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("gorouting running...")
			}
		}
	}()
	fmt.Println("main routing sleep 3s")
	time.Sleep(time.Second * 3)
	//调用cancel取消子线程
	cancel()
	fmt.Println("main routing call cancel")
	time.Sleep(1 * time.Second)
}

func withDeadline() {

	//
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx Done")
		case <-time.After(1 * time.Second):
			fmt.Println("time up...")
		}
	}

}

func withValue() {

	type ContextKey string

	f := func(ctx context.Context, k ContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := ContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, ContextKey("color"))

}
func main() {
	//定义一个空的ctx
	rootCtx := context.TODO()

	//从根ctx衍生子ctx
	ctx, cancel := context.WithCancel(rootCtx)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("cancel called, gorouting 1 exit\n")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("gorouting 1 running...")
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("cancel called, gorouting 2 exit\n")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("gorouting 2 running...")
			}
		}
	}()

	fmt.Println("main routing sleep 3s")
	time.Sleep(time.Second * 3)
	//调用cancel取消子线程
	cancel()
	fmt.Println("main routing call cancel")
	time.Sleep(1 * time.Second)

}
