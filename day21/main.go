package main

import (
	"fmt"
	"math"
)

func sieveOfEratosthenes(value int) []bool {
	f := make([]bool, value)
	for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
		if !f[i] {
			for j := i * i; j < value; j += i {
				f[j] = true
			}
		}
	}
	return f
}

func abundant(i int) bool {
	for div1, sum := 2, 0; div1 < int(math.Sqrt(float64(i))); div1++ {
		if i%div1 == 0 {
			if div2 := (i / div1); i == div2 {
				sum += div1 + 1
			} else {
				sum += (div1 + div2) + 1
			}
			if sum >= i {
				return true
			}
		}
	}
	return false
}

var notPrimes []bool

func main() {
	N := 10000000
	notPrimes = sieveOfEratosthenes(N + 1)
	finalSum := 0
	for x := 10; x <= N; x++ {
		if !notPrimes[x-1] && !notPrimes[x+1] && abundant(x) {
			finalSum += x
		}
	}
	fmt.Printf("%d\n", finalSum)
}
