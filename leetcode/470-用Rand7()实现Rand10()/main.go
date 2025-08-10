package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rand7() int {
	return rand.Intn(7) + 1
}

func rand10() int {
	for {
		num := (rand7()-1)*7 + rand7()
		if num <= 40 {
			return 1 + (num-1)%10
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // 设置随机种子

	counts := make([]int, 11) // 用于验证分布是否均匀

	// 调用多次观察结果分布
	for i := 0; i < 1000000; i++ {
		n := rand10()
		counts[n]++
	}

	// 打印分布
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d: %d\n", i, counts[i])
	}
}
