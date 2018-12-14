package main

import (
	"fmt"
	"math"
)

func sumTo(arr []int, target int) bool {
	var cnt int
	mp := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		k := target - arr[i]
		if _, ok := mp[k]; ok {
			cnt++
			if cnt > 2 {
				return false
			}
		}
		mp[arr[i]] = arr[i]
	}
	return cnt == 1
}

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
	var cnt, sum int
	sieve := sieveOfEratosthenes(6300)
	nums := []int{1, 3}
Loop:
	for n := nums[len(nums)-1] + 1; ; n++ {
		if sumTo(nums, n) {
			nums = append(nums, n)
			cnt, sum = 0, 0
			for _, n := range nums {
				for _, p := range sieve {
					if p == n {
						cnt++
						sum += n
						if cnt == 100 {
							fmt.Println(sum)
							break Loop
						}
						break
					}
				}
			}
		}
	}
}
