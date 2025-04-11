package main

import (
	"bufio"
	"os"
	"fmt"
)

const (
    empty = iota
    O
    X
)

func hasWinner(res map[int]int, size int) bool {
	if res[X] == X * size || res[O] == O * size {
		return true
	}
	return false
}

func canWinX (res map[int]int, size int) bool {
	if res[X] == X * (size - 1) {
		return true
	}
	return false
}

func getResult(board [][]int, n int, m int, k int) bool {

	result := false

	// горизонтали
	for i := 0; i < n; i++ {
		for j := 0; j < m - k + 1; j++ {
			h := make(map[int]int)
			for c := 0; c < k; c++ {
				el := board[i][j+c]
				h[el] += el
			}
			if hasWinner(h, k) { return false }
			if !result && canWinX(h, k) { result = true }
		}
	}

	// вертикали
	for j := 0; j < m; j++ {
		for i := 0; i < n - k + 1; i++ {
			v := make(map[int]int)
			for c := 0; c < k; c++ {
				el := board[i+c][j]
				v[el] += el
			}
			if hasWinner(v, k) { return false }
			if !result && canWinX(v, k) { result = true }
		}
	}

	// диагонали
	for i := 0; i < n - k + 1; i++ {
		for j := 0; j < m - k + 1; j++ {
			d1 := make(map[int]int)
			d2 := make(map[int]int)
			for c := 0; c < k; c++ {
				el1 := board[i+c][j+c]
				el2 := board[i+c][m-j-c-1]
				d1[el1] += el1
				d2[el2] += el2
			}
			if hasWinner(d1, k) { return false }
			if hasWinner(d2, k) { return false }
			if !result && canWinX(d1, k) { result = true }
			if !result && canWinX(d2, k) { result = true }
		}
	}

	return result
}


func main() {
	var in *bufio.Reader
	var out *bufio.Writer

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	consts := map[rune]int{'X': X, 'O': O, '.': empty}

	var t uint16

	fmt.Fscan(in, &t)

	for range t {
		var k int
		fmt.Fscan(in, &k)

		var n, m int
		fmt.Fscan(in, &n, &m)

		board := make([][]int, n)

		for i := range n {
			var s string
			fmt.Fscan(in, &s)
			runes := []rune(s)
			row := make([]int, m)

			for j := range m {
				row[j] = consts[runes[j]]
			}
			board[i] = row
		}

		if getResult(board, n, m, k) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}