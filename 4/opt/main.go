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

// Оптимизация: O(n*m)

func getResult(board [][]int, n int, m int, k int) bool {
	result := false

	checkWin := func(i, j, ci, cj int) bool {
		state := make([]int, 3)
	
		for 0 <= i && i < n && 0 <= j && j < m {
			state[board[i][j]]++
			i += ci
			j += cj
			// скользящее окно
			if 0 <= i - k * ci && i - k * ci < n && 0 <= j - k * cj && j - k * cj < m {
				if state[X] == k || state[O] == k {
					return true
				}
				if !result && state[X] == k - 1 && state[empty] == 1 {
					result = true
				}
				state[board[i-k*ci][j-k*cj]]--
			}
		}
		return false
	}

	// горизонталь -> right
	for i := 0; i < n; i++ {
		if win := checkWin(i, 0, 0, 1); win {
			return !win
		}
	}

	// вертикаль -> down
	for j := 0; j < m; j++ {
		if win := checkWin(0, j, 1, 0); win {
			return !win
		}
	}
	
	// главная диагональ -> left down
	for i := 0; i < n; i++ {
		if win := checkWin(i, 0, 1, 1); win {
			return !win
		}
	}

	for j := 0; j < m; j++ {
		if win := checkWin(0, j, 1, 1); win {
			return !win
		}
	}

	// побочная диагональ -> right down
	for j := 0; j < m; j++ {
		if win := checkWin(0, j, 1, -1); win {
			return !win
		}
	}

	for i := 0; i < n; i++ {
		if win := checkWin(i, m-1, 1, -1); win {
			return !win
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

		ans := map[bool]string{true: "YES", false: "NO"}
		fmt.Println(ans[getResult(board, n, m, k)])
	}
}