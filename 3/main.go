package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func getResult(data []string) int {
	var b strings.Builder

	even := make(map[string]int)
	odd := make(map[string]int)
	strs := make(map[string]int)

	res := 0
	
	for _, s := range data {
		b.Reset()
		b.Grow((len(s) + 1) / 2 + 1)
		for j := 0; j < len(s); j += 2 {
			b.WriteByte(s[j])
		}

		ev := b.String()
		res += even[ev]
		even[ev]++

		if len(s) == 1 {
			continue
		}

		b.Reset()
		b.Grow((len(s) + 1) / 2 + 1)
		for j := 1; j < len(s); j += 2 {
			b.WriteByte(s[j])
		}

		od := b.String()
		res += odd[od]
		odd[od]++

		// если и четные, и нечетные символы совпали > происходит удваивание res
		res -= strs[s] // поэтому отнимаем кол-во дубликатов слов
		strs[s]++
	}

	return res
}


func main() {
	var in *bufio.Reader
	var out *bufio.Writer

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t uint8

	fmt.Fscan(in, &t)

	for range t {
		var n int
		fmt.Fscan(in, &n)

		strs := make([]string, n)

		for j := 0; j < n; j++ {
			fmt.Fscan(in, &strs[j])
		}

		fmt.Println(getResult(strs))
	}
}