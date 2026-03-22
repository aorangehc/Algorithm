package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	ch := make(chan struct{})

// 	go func() {
// 		for {
// 			<-ch
// 			fmt.Println("0")
// 			ch <- struct{}{}
// 		}
// 	}()

// 	go func() {
// 		for {
// 			<-ch
// 			fmt.Println("1")
// 			ch <- struct{}{}
// 		}
// 	}()

// 	ch <- struct{}{}

// 	time.Sleep(10 * time.Millisecond)
// }

func main() {
	ch := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)

	// 打印 0 的协程
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("0")
			ch <- struct{}{}
			<-ch
		}
	}()

	// 打印 1 的协程
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-ch
			fmt.Println("1")
			ch <- struct{}{}
		}
	}()

	wg.Wait()
	close(ch)
}
