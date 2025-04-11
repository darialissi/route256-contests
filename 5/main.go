package main

import (
	"bufio"
	"os"
	"fmt"
	"slices"
)

func getResult(m int, paints [][2]int, n int, boxes [][2]int) int {
	cnt := 0

	for j := m - 1; j >= 0; j-- {
		box := [2]int{0, 0} // коробка макс. размера для текущей картины
		for i := n - 1; i >= 0; i-- {
			if paints[j][0] > boxes[i][0] {
				break
			}
			if paints[j][1] <= boxes[i][1] {
				box = boxes[i]
			}
		}
		if box[0] == 0 { return -1 } // нет подходящей коробки

		for c := j - 1; c >= 0; c-- { // кладем все подходящие картины в текущую коробку
			if paints[c][0] > box[0] || paints[c][1] > box[1] {
				break
			}
			j-- // сокращаем кол-во проверяемых картин
		}
		cnt++
	}

	return cnt
}


func main() {
	var in *bufio.Reader
	var out *bufio.Writer

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t uint16

	fmt.Fscan(in, &t)

	for range t {
		var n int
		fmt.Fscan(in, &n)

		boxes := make([][2]int, n)

		for i := range n {
			var a, b int
			fmt.Fscan(in, &a, &b)

			if a <= b { 
				boxes[i] = [2]int{a, b}
			} else {
				boxes[i] = [2]int{b, a}
			}
		}

		var m int
		fmt.Fscan(in, &m)
	
		paints := make([][2]int, m)

		for j := range m {
			var c, d int
			fmt.Fscan(in, &c, &d)

			if c <= d {
				paints[j] = [2]int{c, d}
			} else {
				paints[j] = [2]int{d, c}
			}
		}

		slices.SortFunc(boxes, func(m1, m2 [2]int) int {
			return m1[0] - m2[0]
		})

		slices.SortFunc(paints, func(m1, m2 [2]int) int {
			return m1[0] - m2[0]
		})

		fmt.Println(getResult(m, paints, n, boxes))
	}
}