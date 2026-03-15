package main

import (
	"fmt"
	"sync"
)

func main() {
	ss := make([]int, 100)

	for i := 0; i < 100; i++ {
		ss[i] = i
	}

	var wg sync.WaitGroup

	chmap := make(map[int]chan int)
	s := make(chan struct{})

	for i := 0; i < 10; i++ {
		chmap[i] = make(chan int)
		wg.Add(1)

		go func(idx int) {
			defer wg.Done()

			for val := range chmap[idx] {
				fmt.Printf("go id: %d, val : %d \n", idx, val)

				s <- struct{}{}
			}
		}(i)
	}

	for _, i := range ss {
		id := i % 10
		chmap[id] <- i
		<-s
	}

	for i, _ := range chmap {
		close(chmap[i])

		delete(chmap, i)
	}

	wg.Wait()

	close(s)

	fmt.Println("finish")
}
