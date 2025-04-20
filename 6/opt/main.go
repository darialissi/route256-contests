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
	curSrs := make([]int, m)
	procImg := make([]int, m)
	covImg := 0

	// метод двух указателей O(n)
	for l, r := 0, 0; r < len(data); r++ { // покрывает все изображения
		if procImg[data[r].image] == 0 {
			covImg++
		}
		procImg[data[r].image]++
		curSrs[data[r].image] = data[r].server

		for l <= r && covImg == m { // сужает диапазон
			if curDif := data[r].time - data[l].time; curDif < dif {
				dif = curDif
				copy(srs, curSrs)
			}
			procImg[data[l].image]--
			if procImg[data[l].image] == 0 {
				covImg--
			}
			l++
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