package main

import (
	"bufio"
	"fmt"
	"os"
)

func result(s string) string {
	t := []byte{}
	p := 0

	n := len(s)

	for i := 0; i < n; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			if p == 0 {
				p = int(s[i] - '0')
			} else {
				p = p*10 + int(s[i]-'0')
			}
		} else {
			// 先左移
			m := len(t)
			if m != 0 {
				j := p % m
				tmp := []byte{}
				tmp = append(tmp, t[j:]...)
				tmp = append(tmp, t[:j]...)

				t = tmp
			}

			p = 0

			// 再修改
			if s[i] == 'R' {
				for j := 0; j < len(t)/2; j++ {
					t[j], t[len(t)-1-j] = t[len(t)-1-j], t[j]
				}
			} else {
				t = append(t, s[i])
			}
		}
	}

	return string(t)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n int

	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(reader, &s)

		ans := result(s)

		fmt.Fprintln(writer, ans)
	}

}
