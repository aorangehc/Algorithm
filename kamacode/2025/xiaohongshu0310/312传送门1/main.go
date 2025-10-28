package main

// import "fmt"

// func main() {
// 	var n int
// 	fmt.Scanf("%d", &n)

// 	ans := 0
// 	for i := 0; i < n; i++ {
// 		var num int
// 		fmt.Scanf("%d", &num)

// 		if num < 0 {
// 			ans -= num
// 		} else {
// 			ans += num
// 		}
// 	}

// 	fmt.Println(ans)
// }
import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	const bufferSize = 1 << 20
	reader := bufio.NewReaderSize(os.Stdin, bufferSize)
	writer := bufio.NewWriterSize(os.Stdout, bufferSize)

	defer writer.Flush()

	var n int

	fmt.Fscan(reader, &n)

	ans := 0

	for i := 0; i < n; i++ {
		var num int

		fmt.Fscan(reader, &num)

		if num < 0 {
			ans -= num
		} else {
			ans += num
		}
	}

	fmt.Fprintln(writer, ans)
}