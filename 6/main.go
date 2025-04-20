package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
	"sort"
	"strconv"
)

type Info struct {
	server int
	image int
	time int
}

func getResult(n int, th []int, m int, wg []int) (int, []int) {
	data := []Info{}

	for i := range n {
		for j := range m {
			data = append(data, Info{
				server: i + 1,
				image: j,
				time: (wg[j] + th[i] - 1) / th[i],
			})
		}
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].time < data[j].time
	})

	srs := make([]int, m) // итоговый список серверов
	dif := math.MaxInt32

	// ищем отрезки, покрывающие все изображения
	for l, r := 0, len(data); l < r - m + 1; l++ {
		curSrs := make([]int, m)
		covImg := 0
		for i := l; i < r; i++ {
			if curSrs[data[i].image] == 0 {
				covImg++
			}
			curSrs[data[i].image] = data[i].server

			if covImg == m {
				if curDif := data[i].time - data[l].time; curDif < dif {
					dif = curDif
					copy(srs, curSrs)
				}
				break
			}
		}
	}

	return dif, srs
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

		th := make([]int, n)
		for i := range n {
			var s string
			fmt.Fscan(in, &s)
			v, _ := strconv.Atoi(s)
			th[i] = v
		}

		var m int
		fmt.Fscan(in, &m)

		wg := make([]int, m)
		for i := range m {
			var s string
			fmt.Fscan(in, &s)
			v, _ := strconv.Atoi(s)
			wg[i] = v
		}

		dif, srs := getResult(n, th, m, wg)

		fmt.Println(dif)

		for i := range srs {
			fmt.Printf("%d ", srs[i])
		}
	}
}