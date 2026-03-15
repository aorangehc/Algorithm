// 同时开启 100 个协程，分成 1-100 号协程
// 1 号协程打印尾数为 1 的数字，2 号协程只打印尾号为 2 的数字
// 以此类推，顺序打印 1-1000 整数以及其对应的协程号

package main

import (
	"fmt"
	"time"
)

func main() {
	s := make(chan struct{})
	// 通过 map 的 key保证协程的顺序
	m := make(map[int]chan int, 100)
	// 填充 map 初始化 channel
	for i := 1; i <= 100; i++ {
		m[i] = make(chan int)
	}
	// 开启协程，死循环打印
	for i := 1; i <= 100; i++ {
		go func(id int) {
			for {
				num := <-m[id]
				fmt.Println(num, id)
				s <- struct{}{}
			}
		}(i)
	}

	// 循环 1-1000，并把值传递给匹配的 map
	// 然后通过s限制循环打印

	for i := 1; i <= 1000; i++ {
		id := i % 10
		if id == 0 {
			id = 10
		}
		m[id] <- i
		// 通过s这个来控制打印顺序
		// 每次遍历一次i，都通过s阻塞协程的打印，最后打印完毕
		<-s
	}

	time.Sleep(10 * time.Second)
}
