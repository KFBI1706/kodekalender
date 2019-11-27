package main

import (
	"fmt"
	"io/ioutil"
	"math"
)

func sieveOfEratosthenes(value int) (primes []int) {
	f := make([]bool, value)
	for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
		if !f[i] {
			for j := i * i; j < value; j += i {
				f[j] = true
			}
		}
	}
	for i := 2; i < value; i++ {
		if !f[i] {
			primes = append(primes, i)
		}
	}
	return
}

func main() {
	var LEN, start, stop, sum, max, prime int
	dat, _ := ioutil.ReadFile("./input-palindrom.txt")
	s := string(dat)
	primes := sieveOfEratosthenes(250000)
	LEN = len(s)
	for i := 0; i < LEN; i += 2 {
		for j := 0; j < 2; j++ {
			sum = int(s[i]-'0') * (j + 1)
			for start, stop = i-2+j*2, i+2; start >= 0 && stop < LEN && s[start] == s[stop]; start, stop = start-2, stop+2 {
				sum += int(s[start]-'0') * 2
				if sum > max {
					for _, prime = range primes {
						if prime == sum {
							max = sum
							break
						}
					}
				}
			}
		}
	}
	fmt.Println(max)
}
