package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
)

type currency uint16

const (
    RUB currency = iota
    USD currency = iota
    EUR currency = iota
)

type exchange struct {
	from currency
	to currency
}

type ExchangeRate struct {
	exc exchange
	rate float64
}

func getRateByCur(banks [][]ExchangeRate, from_cur currency, to_cur currency) map[int]float64 {
	data := make(map[int]float64) // bank: rate

	for i := range banks {
		for j := range banks[i] {
			if banks[i][j].exc.from == from_cur && banks[i][j].exc.to == to_cur {
				data[i] = banks[i][j].rate
				break
			}
		}
	}

	return data
}

func getResult(banks [][]ExchangeRate) float64 {
	resultUSD := 0.0

	rubEur := getRateByCur(banks, RUB, EUR)
	usdEur := getRateByCur(banks, USD, EUR)
	eurUsd := getRateByCur(banks, EUR, USD)                                                           
	rubUsd := getRateByCur(banks, RUB, USD)
	usdRub := getRateByCur(banks, USD, RUB)
	eurRub := getRateByCur(banks, EUR, RUB)

	for b1 := range rubEur {
		// rub > eur > usd                                
		for b2 := range eurUsd {
			if b1 != b2 {
				resultUSD = math.Max(resultUSD, rubEur[b1] * eurUsd[b2])
			}
		}
		// rub > eur > rub > usd
		for b2 := range eurRub {
			for b3 := range rubUsd {
				if b1 != b2 && b2 != b3 && b1 != b3 {
					resultUSD = math.Max(resultUSD, rubEur[b1] * eurRub[b2] * rubUsd[b3])
				}
			}
		}
	}

	// rub > usd
	for b1 := range rubUsd {
		resultUSD = math.Max(resultUSD, rubUsd[b1])
		// rub > usd > rub > usd
		for b2 := range usdRub {
			for b3 := range rubUsd {
				if b1 != b2 && b2 != b3 && b1 != b3 {
					resultUSD = math.Max(resultUSD, rubUsd[b1] * usdRub[b2] * rubUsd[b3])
				}
			}
		}
		// rub > usd > eur > usd
		for b2 := range usdEur {
			for b3 := range eurUsd {
				if b1 != b2 && b2 != b3 && b1 != b3 {
					resultUSD = math.Max(resultUSD, rubUsd[b1] * usdEur[b2] * eurUsd[b3])
				}
			}
		}
	}

	return resultUSD
}


func main() {
	var in *bufio.Reader
	var out *bufio.Writer

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t uint16

	fmt.Fscan(in, &t)

	inputExc := []exchange{
		{ from: RUB, to: USD },
		{ from: RUB, to: EUR },
		{ from: USD, to: RUB },
		{ from: USD, to: EUR },
		{ from: EUR, to: RUB },
		{ from: EUR, to: USD },
	}

	for range t {

		banks := make([][]ExchangeRate, 3)

		for i := range 3 {
			banks[i] = make([]ExchangeRate, 6)

			for j := range 6 {
				var n, m uint8
				fmt.Fscan(in, &n, &m)
				banks[i][j] = ExchangeRate{ inputExc[j], float64(m) / float64(n) }
			}
		}

		fmt.Fprintln(out, getResult(banks))
	}
}