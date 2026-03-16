package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Print("a1")
			ch1 <- struct{}{}
			<-ch1
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-ch1
			fmt.Print("b2")
			ch2 <- struct{}{}

			<-ch2
			ch1 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-ch2
			fmt.Println("c3")
			ch2 <- struct{}{}
		}
	}()

	wg.Wait()
	close(ch1)
	close(ch2)
}
