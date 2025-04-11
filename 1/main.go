package main

import (
	"bufio"
	"fmt"
	"os"
)

func isValid(s string) bool {
	runes := []rune(s)

	orig := runes[0]
	c := 0

	for i := 1; i < len(runes); i++ {
		if c > 1 {
			return false
		}
		if runes[i] != orig {
			c++
		} else if runes[i] == orig && c == 1 {
			c--
		}
	}

	return c == 0
}


func main() {
	var in *bufio.Reader
	var out *bufio.Writer

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int

	fmt.Fscan(in, &n)

	for range n {
		var s string
		fmt.Fscan(in, &s)

		if isValid(s) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}