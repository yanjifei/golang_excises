package main

import (
	"fmt"
	"math"
	"math/big"
)

const (
	maxPrec   = 100
	iteration = 10000
)

func main() {
	n := 8
	for {
		fmt.Println("How many fractional digits do you want to generate?")
		fmt.Scanf("%d", &n)
		if n >= 0 && n <= 100 {
			break
		} else {
			fmt.Printf("Limits : %d digits\nTry again.\n", maxPrec)
		}
	}

	chl := make(chan *big.Float, iteration)

	for i := 0; i < iteration; i++ {
		go calcPi(chl, float64(i))
	}
	result := new(big.Float).SetPrec(maxPrec)
	for i := 0; i < iteration; i++ {
		result.Add(result, <-chl)
	}

	fmt.Printf("result = %1.*f\n", n, result)
	//fmt.Println(pi(iteration))

}

//Calculate Pi
//https://en.wikipedia.org/wiki/Leibniz_formula_for_%CF%80
func calcPi(chl chan *big.Float, i float64) {
	minusOnePow := new(big.Float).SetFloat64(math.Pow(-1, i))
	four := new(big.Float).SetInt64(4)
	twoMulIPlusOne := new(big.Float).SetFloat64(2*i + 1)
	div := new(big.Float).SetPrec(maxPrec)
	result := new(big.Float).SetPrec(maxPrec)
	div.Quo(minusOnePow, twoMulIPlusOne)
	result.Mul(four, div)
	chl <- result

}

func pi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k <= n; k++ {
		go term(ch, float64(k))
	}
	f := 0.0
	for k := 0; k <= n; k++ {
		f += <-ch
	}
	return f
}

func term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}
